# -*- coding: utf-8 -*-
import sys
import gi

gi.require_version('Gtk', '3.0')
from gi.repository import Gio,Gtk,Gdk


class ImgButton(Gtk.Button):

	def __init__(self, label="",image=None):
		Gtk.Button.__init__(self)
		grid = Gtk.Grid()
		self.add(grid)
		label = Gtk.Label(label)
		grid.attach(label, 2, 1, 1, 1)
		if image:
			grid.attach(image, 0, 1, 1, 1)