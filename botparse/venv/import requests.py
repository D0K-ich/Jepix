import requests
import bs4


url = 'https://vkusvill.ru/goods/sladosti-i-deserty/'
headers = {'User-agent' : 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36'}

def main(url):
    request = requests.get(url, headers)
    return bs4.beautifulsoup(request.text, 'html.parser')

    categories.findall('span', class_='Price__value' )