import requests
import json
import XenAPI
import socket 
import re
# Steps:
#   1.Get list of master nodes 
#   2.For every master node:
#       3. Create session for remote XenAPI
#       4. Get list ov vms from PIM
#       5. Get list of vms from XenAPI

class GetPimVms(object): # Create a new class;  

    def get_list_master_pool(filepath):
        with open(filepath) as fp:
            master_nodes = fp.read().splitliens()
        return master_nodes


    # Step 2.Connect to remote xenapi    
    def create_session(url, username, password):
        try:
            session = XenAPI.Session(url, ignore_ssl=True)
            print session
            session.xenapi.login_with_password(username, password)
        except Exception, e:
            if hasattr(e, 'details') and e.details[0] == 'HOST_IS_SLAVE':
                # Redirect to cluster master
                url = urlparse(url).scheme + '://' + e.details[1]
                print url

                try:
                    session = XenAPI.Session(url, ignore_ssl=True)
                    session.login_with_password(username, password)
                except Exception, e:
                    handle_login_error(e)
            else:
                handle_login_error(e)
        return (session)

    def _list_of_vms_pim(self): # Step 3.List of vms from PIM
        vms = list()
        return_list = list()
        sub_part_host = self.poolname.split(".")[0].split("-")[1] # Get part of the hostname ex: xen1
        sub_part_host_res = re.sub(r"\d+","*",sub_part_host) # Substitute all numbers in sub_part_host to * ex: xen*
        regexp = re.sub(sub_part_host,sub_part_host_res,self.poolname) # Make regexp for pim
        query = "http://your_cloudstack_server/v1/asset?keyword={}&assetTypeId=2&deleted=false&highlight=false&aggregation=false&resultsPerPage=10000".format(regexp) # Change value to poolname
        rdata = requests.get(query)
        rdata = rdata.json()
        for items in rdata['items']:
            for child_links in items['links']['children']:
                vms.append(child_links) # need to create a list,as child_links is a string of links to child objects and I can't work with them directly 
            for vmname in vms:
                if vmname['assetTypeId'] == 1 and vmname['displayName'] not in return_list:
                    return_list.append(vmname['displayName'])
        for name in return_list:
            print(name)
    
    
    def _list_of_vms_hv(self,session): # Step 4.List of vms from hypervisor
        vmArray = []
        vms = session.xenapi.VM.get_all() #Get all the vms running on the xen hypervisor
        # Loop through each VM in vms, we'll skip it if the vm is a template, control domain, or a snapshot
        for vm in vms:
                if session.xenapi.VM.get_is_a_template(vm):
                        continue
                elif session.xenapi.VM.get_is_control_domain(vm):
                        continue
                elif session.xenapi.VM.get_is_a_snapshot(vm):
                        continue
                else:
                        if session.xenapi.VM.get_name_label(vm) not in vmArray:
                                vmArray.append(session.xenapi.VM.get_name_label(vm))


    def comp(self,list1, list2): # Step 5.Compare lists 
        difference = list()
        for val in list1: 
            if val in list2: 
                continue
            else: 
                difference.append(val)
        print(difference)
        # check alternative difference.append(set(list1) & set(list2))
                  


OBJECTS = GetPimVms(create_session(""))
OBJECTS.comp(OBJECTS._list_of_vms_pim(),OBJECTS._list_of_vms_hv())