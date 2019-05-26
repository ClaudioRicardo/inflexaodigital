# -*- coding: utf-8 -*-
# author:Claudio Azevedo (claudwolf@hotmail.com) 
import requests
from bs4 import BeautifulSoup


def get_data(url):
    try:
        result = requests.get(url.strip())
        if result.status_code != 200:
           result  = False

    except requests.exceptions.ConnectionError as r:
        print("ERROR: " + str(r))
        return  False

    return result




def get_table(url):
    data_html = get_data(url)
    data_three = BeautifulSoup(data_html.text, "html.parser")
    data_three.prettify()
    table = data_three.find("table",{"class": "table"})
    data_table_list = []

    ind = 0
    for line in table.find_all("tr"):
        data_table_list.append([])
        for cell in line.find_all({"td","th"}):
            data_table_list[ind].append(cell.text)
        ind+=1

    return data_table_list


if __name__=="__main__":
    url = "http://127.0.0.1:8000/blog/a-coleta-de-dados-como-princ%C3%ADpio-da-minera%C3%A7%C3%A3o-de-dados-na-web-com-python/"
    print(get_table(url))



