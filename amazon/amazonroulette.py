import urllib
from xml.dom import minidom
import os.path
from collections import defaultdict
import random
import sys

class bcolors:
    HEADER = '\033[95m'
    OKBLUE = '\033[94m'
    OKGREEN = '\033[92m'
    WARNING = '\033[93m'
    FAIL = '\033[91m'
    ENDC = '\033[0m'
    BOLD = '\033[1m'
    UNDERLINE = '\033[4m'


if len(sys.argv) == 1:
    print "python amazonroulette.py [file name with buckets names]"
    sys.exit()


file = sys.argv[1]


with open(file) as f:
    for line in f:
        line = line.rstrip()
        xml_str = urllib.urlopen("http://" + line).read()
        extensions = defaultdict(list)
        xmldoc = minidom.parseString(xml_str)
        obs_values = xmldoc.getElementsByTagName('Key')

        for i in obs_values:
            path = i.firstChild.data
            extension = os.path.splitext(i.firstChild.data)[1][1:]
            if extensions.has_key(extension):
                extensions[extension].append(path)
            else:
                extensions[extension] = [path]

        helper = 0
        while helper == 0:
            print "Founded extensions\n " + line
            for i in (extensions):
                print i, len(extensions[i])
            choosen_ext = raw_input("Type extension or go to [N]ext target\n> ")
            while 1:
                if choosen_ext in extensions:
                    for i in extensions[choosen_ext]:
                        print i
                    decision = raw_input("Check All[A] Roulette[R] [B]ack \n> ")
                    if decision == "A":
                        for i in extensions[choosen_ext]:
                            req = urllib.urlopen("http://" + line + "/" + i)
                            if req.code == 200:
                                print bcolors.OKGREEN + i + " -----> " + str(req.code) + bcolors.ENDC + "\n"
                            else:
                                print bcolors.FAIL + i + " -----> " + str(req.code) + bcolors.ENDC + "\n"

                    elif decision == "R":
                        ran = random.randint(0,len(extensions[choosen_ext]) -1 )
                        url = "http://" + line + "/" + extensions[choosen_ext][ran]
                        req = urllib.urlopen(url)
                        if req.code == 200:
                            print bcolors.OKGREEN + url + " -----> " + str(req.code) + bcolors.ENDC

                        else:
                            print bcolors.FAIL + url + " -----> " + str(req.code) + bcolors.ENDC
                            continue
                    elif decision== "B":
                        break
                    else:
                        print "Wrong choice1\n "
                        continue
                if choosen_ext == "N":
                    helper = helper + 1
                    break
                else:
                    print "Wrong choice"
                    break
