import XenAPI
import re
import requests
import sys
from ruamel.yaml import YAML

class GetPimVms(object): 
    def __init__(self,hostname,logfiledescr):
	    self.hostname = hostname
        self.logfiledescr = logfiledescr  
    # Create session to connect to remote XenAPI
	def create_session(self,url, username, password):
        session = XenAPI.Session(url)
        session.xenapi.login_with_password(username, password)
        return session

	#List of vms from PIM
    def _list_of_vms_pim(self):
            vms = list()
        return_list = list()
        sub_part_host = self.hostname.split(".")[0].split("-")[1] # Get part of the hostname ex: xen1
        sub_part_host_res = re.sub(r"\d+","*",sub_part_host) # Substitute all numbers in sub_part_host to * ex: xen*
        regexp = re.sub(sub_part_host,sub_part_host_res,self.hostname) # Make regexp for pim
        query = "http://(HOSTNAME)/(X)/v1/asset?keyword={}&assetTypeId=2&deleted=false&highlight=false&aggregation=false&resultsPerPage=10000".format(regexp) # Change value to poolname
        rdata = requests.get(query)
        rdata = rdata.json()
        for items in rdata['items']:
        for child_links in items['links']['children']:
            vms.append(child_links) # need to create a list,as child_links is a string of links to child objects and I can't work with them directly 
        for vmname in vms:
            if vmname['assetTypeId'] == 1 and vmname['displayName'] not in return_list:
            return_list.append(str(vmname['displayName'])) # need to convert to str as default type of vmname['displayName'] is unicode
        return return_list
        
    def _list_of_vms_hv(self,session): #List of vms from hypervisor
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
        return vmArray
        #!!!!!!!!!!!!!!!!! Need to close the session !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!#

    def comp(self,list1, list2): #Compare lists 
	    diff = list(set(list1).symmetric_difference(set(list2))) # Values that are not in both sets
	    if len(diff) != 0:
          self.notif(diff)
        
        # Write difference in result file 
    def notif(self,diff):
        for i in diff:
	        output['diff'] += "xapi_pim_diff{master=\"" + self.hostname + "\",vm=\""+ i + "\"} 1\n"  # 1 means that

with open(sys.argv[1],'r') as config:  # Read data from config file
 yaml = YAML()
 data = yaml.load(config)
output = {
  "diff": "# HELP xapi_pim_diff Difference in the info provided by XenAPI and PIM.\n# TYPE xapi_pim_diff  gauge\n",
  "auth_fail": "# HELP xen_authentication_fail Failed to authenticate to XenAPI.\n# TYPE xen_authentication_fail gauge\n"
}
logfiledescr = open(data['logfile'],'w+') # File is opened here because it should be overwritten every time when script executes
masters = list()
query = "http://(HOSTNAME)/(X)/v1/asset?keyword=(TEMPLATEE)&assetTypeId=2&deleted=false&highlight=false&aggregation=false&resultsPerPage=10000" 
rdata = requests.get(query)
rdata = rdata.json()
for i in rdata['items']: # Here we gather all pool that fit the pattern pool*-xen1.sys.ams1.cloudsys.tmcs. They will be used for comparing
 masters.append(i['doc']['Name'])
for master in masters: 
 url = "http://" + str(master) # Create url string to connect to xenapi
 OBJECT = GetPimVms(master,logfiledescr) # as a second param file descriptor is passed to write result file
 try:
  session = OBJECT.create_session(url,data['user'],data['password'])
 except:
  output['auth_fail'] += 'xen_pim_authentication_fail{pool="%s"} 1\n'%(master)
  continue
 OBJECT.comp(OBJECT._list_of_vms_pim(),OBJECT._list_of_vms_hv(session)) # compare PIM list and Xen list
for value in output:
 logfiledescr.write(output[value])
logfiledescr.close() # Close file
