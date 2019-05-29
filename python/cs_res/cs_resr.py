#!/usr/bin/env python
from cs import CloudStack, read_config
import json
import sys

class load_resources(object):
    def __init__(self, tenant,cs):
        self.tenant = tenant
        self.limits_used = {}
        self.limits_total = {}
        self.cs = cs

    def get_resources(self):
        self.projectid=self.cs.listProjects(name=self.tenant)['project'][0]['id']
        limits = self.cs.listProjects(name=self.tenant)['project'][0]
        self.limits_used = {"memory": limits['memorytotal'],"cpu": limits['cputotal'],"primarystorage": limits['primarystoragetotal'],"volume": limits["volumetotal"],"secondarystorage": limits['secondarystoragetotal']}
        self.limits_total = {"volume": limits['volumelimit'],"primarystorage": limits['primarystoragelimit'],"cpu": limits['cpulimit'],"memory": limits['memorylimit'],"secondarystorage": limits['secondarystoragelimit']}

        
    def write_to_prom(self):
        with open('res.prom',"a") as prom_log:
            for res in self.limits_total:
                logstring_used = "cloudstack_%s_used{project=%s} %s\n"%(res,self.tenant,self.limits_used.get(res))
                logstring_total = "cloudstack_%s_total{project=%s} %s\n"%(res,self.tenant,self.limits_total.get(res))
                prom_log.write(logstring_used)
                prom_log.write(logstring_total)



endpoint = 'http://endpoint'
apikey = 'your apikey'
secretkey = 'your secretkey'
cs = CloudStack(endpoint=endpoint,key=apikey,secret=secretkey)
projects_list = list()
for project in cs.listProjects()['project']:
    projects_list.append(project['displaytext'])
for project in projects_list:
    RESOURCES = load_resources(project,cs)
    RESOURCES.get_resources()
    RESOURCES.write_to_prom()
