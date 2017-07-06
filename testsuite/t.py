import requests
from bs4 import BeautifulSoup
import argparse, subprocess

parser = argparse.ArgumentParser()
parser.add_argument("-i", dest="url", help="url value")

args = parser.parse_args()
URL = args.url
download_list = open('DownloadList', 'w')

response = requests.get(url=URL)
response.raise_for_status()

soup = BeautifulSoup(response.content, "lxml")

for link in soup.find_all('a'):
    href = link['href']
    process = subprocess.Popen('ls | grep %s' % href, shell=True, stdout=subprocess.PIPE)
    out, err = process.communicate()
    out = out.decode('utf-8')
    if href in out:
        continue
    if href != '../':
        download_list.write(URL+href)
        download_list.write('\n')

download_list.close()
subprocess.Popen("tmux new-session -d -s downloader 'wget -i DownloadList; bash -i'", shell=True)
