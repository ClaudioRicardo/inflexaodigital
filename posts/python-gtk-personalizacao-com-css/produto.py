 # -*- coding: utf-8 -*-
import gi
gi.require_version('Gtk', '3.0')
from gi.repository import Gtk
from imgbutton import ImgButton

class Produto():
	
	def __init__(self):
		self.grid = Gtk.Grid(column_homogeneous=True)
		self.grid.set_name("grd_form_prod")

		btn_salvar = Gtk.Button(label="Salvar")  #Gtk.Button(label="Compras")	
		btn_pesquisar = Gtk.Button(label="Pesquisar")			
		
		lbl_produto = Gtk.Label("Produto")
		lbl_produto.set_name("title")

		lbl_cod_prod = Gtk.Label("Código")
		lbl_cod_prod.set_name("lbl_fields")
		lbl_cod_prod.set_halign(Gtk.Align.START)

		inp_cod_prod = Gtk.Entry()
		inp_cod_prod.set_editable(False)
		
		lbl_nome_prod = Gtk.Label("Nome")
		lbl_nome_prod.set_name("lbl_fields")
		lbl_nome_prod.set_halign(Gtk.Align.START)
		inp_nome_prod = Gtk.Entry()

		lbl_valor_atual = Gtk.Label("R$ Valor atual")
		lbl_valor_atual.set_name("lbl_fields")
		lbl_valor_atual.set_halign(Gtk.Align.START)

		inp_val_atual = Gtk.Entry()

		# hbox1 = Gtk.Box(spacing=6)
		# hbox1.set_name("box_titulo")
		# hbox1.pack_start(lbl_produto, True, True, 0)



		self.grid.attach(lbl_produto,0,0,4,1)
		self.grid.attach(lbl_cod_prod,0,1,1,1)
		self.grid.attach(inp_cod_prod,0,2,1,1)
		self.grid.attach(lbl_nome_prod,1,1,3,1)
		self.grid.attach(inp_nome_prod,1,2,3,1)
		self.grid.attach(lbl_valor_atual,0,3,1,1)
		self.grid.attach(inp_val_atual,0,4,1,1)
		self.grid.set_row_spacing(10)
		self.grid.set_column_spacing(10)
		self.grid.props.margin = 20
		
		self.grid.attach(btn_salvar,3,5,1,1)
		self.grid.attach(btn_pesquisar,2,5,1,1)

		

	def get_form(self):
		return self.grid	


# class Produto(Gtk.Window):
# 	def __init__(self):
# 		Gtk.Window.__init__(self, title="Produto")
# 		self.set_default_size(600,400)	
# 		#self.set_modal(True)
# 		#self.set_border_width(10)
# 		#self.set_resizable(False)

# 		grid = Gtk.Grid(column_homogeneous=True)
# 		self.add(grid)
# 		boxform = Gtk.Box(spacing=6)

# 		label1 = Gtk.Label(" Teste ")
# 		boxform.pack_start(label1, True, True, 0) 

# 		img_prod = Gtk.Image.new_from_file("imgs/171-lab.png")
# 		img_comp = Gtk.Image.new_from_file("imgs/059-cart.png")
# 		img_estoque = Gtk.Image.new_from_file("imgs/369-table.png")
# 		img_esti = Gtk.Image.new_from_file("imgs/156-stats-dots.png")
# 		img_bkg = Gtk.Image.new_from_file("imgs/background.jpg")
		
# 		btn_produtos = ImgButton(label="Produtos",image=img_prod)
# 		#btn_produtos.connect("clicked",self.open_produto)

# 		btn_compras = ImgButton(label="Comprar",image=img_comp)#Gtk.Button(label="Compras")
# 		btn_estoque  = ImgButton(label="Estoque",image=img_estoque)#Gtk.Button(label="Estoque")
# 		btn_estimativas  = ImgButton(label="Estimativas",image=img_esti) #Gtk.Button(label="Estimativas")

# 		grid.attach(btn_produtos,0,1,1,1)
# 		grid.attach(btn_compras,1,1,1,1)
# 		grid.attach(btn_estoque,2,1,1,1)
# 		grid.attach(btn_estimativas,3,1,1,1)
# 		grid.attach(boxform,0,0,1,1)