import requests
import yaml
import sys

class Player():
    def __init__(self,account_name):
        self.account_name = account_name
        self.account_id = 0
        self.global_raiting = 0
        self.battles = 0
        self.avg_blocked = 0
        self.hits_percents = 0
        self.wins = 0
        self.losses = 0
        self.max_frags = 0
        self.tanking_factor = 0
        self.shots = 0
        self.max_damage = 0
        self.trees_cut = 0
        self.avg_xp = 0
        self.avg_assist = 0
        self.draws = 0
        self.damage_received = 0
        self.capture = 0
        


    def get_acc_id(self):
        query = 'https://api.worldoftanks.ru/wot/account/list/?application_id=b9092ffc2fc7c37f97d22e4e485cb74c&search={0}'.format(self.account_name)
        resp = requests.get(query).json()
        self.account_id = resp['data'][0]['account_id']

    def stat_collect(self):
        query = 'https://api.worldoftanks.ru/wot/account/info/?application_id=b9092ffc2fc7c37f97d22e4e485cb74c&account_id={0}'.format(self.account_id)
        resp = requests.get(query).json()
        resp_part = resp['data'][str(self.account_id)]
        resp_all = resp['data'][str(self.account_id)]['statistics']['all']
        # Collecting stat
        self.global_raiting = resp_part['global_rating']
        self.battles = resp_all['spotted']
        self.avg_blocked = resp_all['avg_damage_blocked']
        self.hits_percents = resp_all['hits_percents']
        self.wins = resp_all['wins']
        self.losses = resp_all['losses']
        self.draws = resp_all['draws']
        self.max_frags = resp_all['max_frags']
        self.tanking_factor = resp_all['tanking_factor']
        self.shots = resp_all['shots']
        self.max_damage = resp_all['max_damage']
        self.trees_cut = resp_part['statistics']['trees_cut']
        self.avg_xp = resp_all['battle_avg_xp']
        self.avg_assist = resp_all['avg_damage_assisted']
        self.damage_received = resp_all['damage_received']
        self.capture = resp_all['capture_points']

    def write_for_prom(self,filedescr):
        for val in self.__dict__:
            if val != 'account_name' and val != 'account_id':
                query = "stat_%s{user=%s} %s\n"%(val,self.account_name,self.__dict__[val])
                filedescr.write(query)

        

        
with open(sys.argv[1],"r") as yaml_data:
    data = yaml.safe_load(yaml_data)
filedescr = open(data['logfile'],"w+")
for account in data['accounts']:
    player = Player(account)
    player.get_acc_id()
    player.stat_collect()
    player.write_for_prom(filedescr)
filedescr.close()


