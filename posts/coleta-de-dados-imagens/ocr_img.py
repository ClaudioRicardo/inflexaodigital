# -*- coding: utf-8 -*-
from PIL import Image
import requests
import pytesseract
from io import BytesIO



''' 
	LINUX
	Para que a bateria pytesseract funcione pode ser necessária
	a instalação das libs abaixo

	sudo apt update
	sudo apt install tesseract-ocr
	sudo apt install libtesseract-dev
'''


def get_image():
    '''
       Pega uma imagem na internet para usar como exemplo
    '''
    img_path = 'https://slideplayer.com.br/slide/1723756/7/images/6/LEITURA+E+AN%C3%81LISE+DE+TEXTOS.jpg' #'https://i.ytimg.com/vi/kFVHslQbf6o/hqdefault.jpg'
    response = requests.get(img_path)
    img = Image.open(BytesIO(response.content))
    return img



if __name__=="__main__":
    img = get_image()
    
    '''
		Passa a imagem para o metodo image_to_string que retornará o texto contido na imagem
    '''
    print(pytesseract.image_to_string(img))    