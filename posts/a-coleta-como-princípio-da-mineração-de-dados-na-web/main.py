# -*- coding: utf-8 -*-
# author:Claudio Azevedo (claudwolf@hotmail.com) 
import requests
from bs4 import BeautifulSoup


def get_data(url):
    '''
        Função receberá uma url e com a ajuda do método get da bateria
        requests retornará um documento HTML ou uma mensagem de erro e um
        return False caso o status_code for diferente de 200.
    '''
    try:
        result = requests.get(url.strip())
        if result.status_code != 200:
           result  = False

    except requests.exceptions.ConnectionError as r:
        print("ERROR: " + str(r))
        return  False

    return result




def get_table(url):
    '''
       A função vai trabalhar com a página html coletada pela função
       get_data. Com a bateria BeautifulSoup busca-se uma table
       especifica e extrai-se os dados, alimentando uma matriz com eles.
    '''
    data_html = get_data(url)
    data_three = BeautifulSoup(data_html.text, "html.parser")
    data_three.prettify()

    #Buscando dentro página uma table com a class="table"  

    table = data_three.find("table",{"class": "table"})
    data_table_list = []

    ind = 0
    #Percorrendo a linha
    for line in table.find_all("tr"):
        data_table_list.append([])
        #Percorrendo as células de cada linha
        for cell in line.find_all({"td","th"}):
            #Armazenando o texto das células na matriz
            data_table_list[ind].append(cell.text)
        ind+=1

    return data_table_list


if __name__=="__main__":
    url = "https://gist.githubusercontent.com/ClaudioRicardo/1f04d5c865c719d50ecd5fe3ada58cab/raw/988ad47bba3bf8b54bd69145abdad555cf309863/tabela_alvo.html"
    print(get_table(url))



