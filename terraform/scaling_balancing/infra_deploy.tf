provider "aws" {
    access_key = "${var.AWS_ACCESS_KEY}"
    secret_key = "${var.AWS_SECRET_KEY}"
    region = "${var.AWS_REGION}"
}

resource "aws_key_pair" "deploy_vm" {
    key_name = "deploy_vm"
    public_key = "${file("host_key.pub")}"
}


# Security groups

resource "aws_security_group" "access_to_http" {
    name = "elb_http"
    description = "Allow inbound traffic to port 80"
    ingress {
        from_port = 80
        to_port = 80
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }
    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
}
resource "aws_security_group" "access_to_lb_instances" {
    ingress {
        from_port = 22
        to_port = 22
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }
    # for healthchecks
    ingress {
        from_port = 80
        to_port = 80
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }
    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
  
}




# Load balancer

resource "aws_elb" "simple_balancer" {
    name = "simpleelb"
    cross_zone_load_balancing = false
    security_groups = ["${aws_security_group.access_to_http.id}"]
    availability_zones = ["${var.AWS_REGION}a"]
    listener {
        instance_port = 80
        instance_protocol = "http"
        lb_port = 80
        lb_protocol = "http"
    }
    health_check {
        healthy_threshold = 2
        unhealthy_threshold = 2
        target = "HTTP:80/"
        interval = 30
        timeout = 5
    }
    
}

# Launch configuration
resource "aws_launch_configuration" "launch_config" {
    instance_type = "t2.micro"
    image_id = "ami-0d7511d38c8750092"
    name_prefix = "autoscaling_instance"
    user_data = "${file("scripts/apache2.conf")}"
    key_name = "${aws_key_pair.deploy_vm.key_name}"
    security_groups = ["${aws_security_group.access_to_lb_instances.id}"]  
}

# Autoscalling group
resource "aws_autoscaling_group" "scale_http_server" {
    name = "scale_http_server"
    min_size = 1
    max_size = 2
    launch_configuration = "${aws_launch_configuration.launch_config.name}"
    force_delete = true
    load_balancers = ["${aws_elb.simple_balancer.id}"]
    health_check_type = "ELB" 
    availability_zones = ["${var.AWS_REGION}a"]
}

# Autoscalling policy
resource "aws_autoscaling_policy" "autoscale_http_servers" {
    name = "autoscale_http_servers"
    scaling_adjustment = 1 // add one server
    adjustment_type = "ChangeInCapacity" // add to the number of the existing servers
    cooldown = 100 // number of seconds before policy can be applied next time
    autoscaling_group_name = "${aws_autoscaling_group.scale_http_server.name}" 
}

resource "aws_autoscaling_policy" "autoremove_http_servers" {
    name = "autoremove_http_servers"
    scaling_adjustment = -1 // add one server
    adjustment_type = "ChangeInCapacity" // add to the number of the existing servers
    cooldown = 100 // number of seconds before policy can be applied next time
    autoscaling_group_name = "${aws_autoscaling_group.scale_http_server.id}" 
}

# Cloudwatch CPU metric for increasing
resource "aws_cloudwatch_metric_alarm" "cpu_check_increase" {
    alarm_name = "cpu_check"
    comparison_operator = "GreaterThanOrEqualToThreshold"
    evaluation_periods = "2" // how often to gather stats
    metric_name = "CPUUtilization"
    period = "60" // how often do comparison with threshold
    statistic = "Average"
    threshold = "80"
    namespace = "AWS/EC2"
    dimensions { // where to apply policy
        AutoScalingGroupName = "${aws_autoscaling_group.scale_http_server.name}"
    }
    alarm_actions = ["${aws_autoscaling_policy.autoscale_http_servers.arn}"] // which policy to apply 
}

# Cloudwatch CPU metric for decreasing
resource "aws_cloudwatch_metric_alarm" "cpu_check_decrease" {
    alarm_name = "cpu_check"
    comparison_operator = "LessThanOrEqualToThreshold"
    evaluation_periods = "2"
    metric_name = "CPUUtilization"
    period = "60"
    statistic = "Average"
    threshold = "20"
    namespace = "AWS/EC2"
    dimensions {
        AutoScalingGroupName = "${aws_autoscaling_group.scale_http_server.name}"
    }
    alarm_actions = ["${aws_autoscaling_policy.autoremove_http_servers.arn}"] 
}

output "elb-dns" {
  value = "${aws_elb.simple_balancer.dns_name}"
}









