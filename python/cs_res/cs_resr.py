from cs import CloudStack
import json
import sys
from ruamel.yaml import YAML

class load_resources(object):
    def __init__(self, tenant,cs,logfiledescr):
        self.tenant = tenant
	self.logfiledescr = logfiledescr
        self.limits_all = {}
        self.cs = cs

    def get_resources(self):
        self.projectid=self.cs.listProjects(name=self.tenant)['project'][0]['id']
        limits = self.cs.listProjects(name=self.tenant)['project'][0]
        for limit in limits:
         if limits[limit] == 'Unlimited':
          limits[limit] = 1000000000
        self.limits_all = {"memory_total": limits['memorytotal'],"cpu_total": limits['cputotal'],"primarystorage_total": limits['primarystoragetotal'],"volume_total": limits["volumetotal"],"secondarystorage_total": limits['secondarystoragetotal'],"volume_used": limits['volumelimit'],"primarystorage_used": limits['primarystoragelimit'],"cpu_used": limits['cpulimit'],"memory_used": limits['memorylimit'],"secondarystorage_used": limits['secondarystoragelimit']}

        
    def write_to_prom(self):
        for res in self.limits_all:
         limitsall[res] += "cloudstack_%s{project=\"%s\"} %s\n"%(res,self.tenant,self.limits_all.get(res))
	 


with open(sys.argv[1],'r') as config:  # Read data from config file
    yaml = YAML()
    data = yaml.load(config)
    limitsall = {
"memory_total": "# HELP  cloudstack_memory_total Total amount of allocated memory that can be used by this project\n# TYPE  cloudstack_memory_total gauge\n",
"cpu_total": "# HELP  cloudstack_cpu_total Total number of allocated vCPU  can be used by this project \n# TYPE  cloudstack_cpu_total gauge\n",
"primarystorage_total": "# HELP  cloudstack_primarystorage_total Total space allocated on primary storage\n# TYPE cloudstack_primarystorage_total gauge\n",
"volume_total": "# HELP  cloudstack_volume_total Total number of disk volumes that can be created by this project\n# TYPE  cloudstack_volume_total gauge\n",
"secondarystorage_total": "# HELP  cloudstack_secondarystorage_total Total space allocated on secondary storage\n# TYPE  cloudstack_secondarystorage_total gauge\n",
"memory_used": "# HELP  cloudstack_memory_used Amount of memory used by the project\n# TYPE cloudstack_memory_used gauge\n",
"cpu_used": "# HELP  cloudstack_cpu_used Number of cores used by the project\n# TYPE  cloudstack_cpu_used gauge\n",
"primarystorage_used": "# HELP  cloudstack_primarystorage_used Amount of primarystorage used by th eproject \n# TYPE  cloudstack_primarystorage_used gauge\n",
"volume_used": "# HELP  cloudstack_volume_used Number of disk volumes used by this ptoject\n# TYPE  cloudstack_volume_used gauge\n",
"secondarystorage_used": "# HELP  cloudstack_secondarystorage_used Amount of secondarystorage used by the project\n# TYPE  cloudstack_secondarystorage_used gauge\n"}
logfiledescr = open(data['logfile'],'w+') # File is opened here because it should be overwritten every time when script executes
cs = CloudStack(endpoint=data['endpoint'],key=data['api'],secret=data['secret'])
projects_list = list()
for project in cs.listProjects()['project']: # gather list of tenants
    projects_list.append(project['displaytext'])
for project in projects_list:
    RESOURCES = load_resources(project,cs,logfiledescr)
    RESOURCES.get_resources()
    RESOURCES.write_to_prom()	
for values in limitsall: 
    logfiledescr.write(limitsall[values])
logfiledescr.close()
