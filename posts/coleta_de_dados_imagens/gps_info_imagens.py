# -*- coding: utf-8 -*-
import sys, os, shutil, base64
from PIL import Image
from PIL.ExifTags import TAGS
from PIL.ExifTags import GPSTAGS


def get_image():
    img_path = 'images/20190518_092637.jpg'
    img = Image.open(img_path);
    return img


def get_exif(img):
    img.verify()
    exif = img._getexif() 
    return exif

def get_geotagging(exif):
    if not exif:
        raise ValueError("No EXIF metadata found")

    geotagging = {}
    for (idx, tag) in TAGS.items():
        if tag == 'GPSInfo':
            if idx not in exif:
                raise ValueError("No EXIF geotagging found")

            for (key, val) in GPSTAGS.items():
                if key in exif[idx]:
                    geotagging[val] = exif[idx][key]

    return geotagging


def get_decimal_from_dms(dms, ref):

    degrees = dms[0][0] / dms[0][1]
    minutes = dms[1][0] / dms[1][1] / 60.0
    seconds = dms[2][0] / dms[2][1] / 3600.0

    if ref in ['S', 'W']:
        degrees = -degrees
        minutes = -minutes
        seconds = -seconds

    return round(degrees + minutes + seconds, 5)

def get_coordinates(geotags):
    lat = get_decimal_from_dms(geotags['GPSLatitude'], geotags['GPSLatitudeRef'])
    lon = get_decimal_from_dms(geotags['GPSLongitude'], geotags['GPSLongitudeRef'])

    return (lat,lon)

if __name__=="__main__":
    img = get_image()
    exif = get_exif(img)
    geotags = get_geotagging(exif)
    coords = get_coordinates(geotags)
    
    print(coords)
