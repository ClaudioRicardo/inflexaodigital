# -*- coding: utf-8 -*-
# arquivo app.py
import sys
import gi

gi.require_version('Gtk', '3.0')
from gi.repository import Gio,Gtk,Gdk


class AppWindow(Gtk.ApplicationWindow):
    '''
        Classe responsável pela criação da janela.
        Essa é a estrutura mais básica de uma janela Python GTK.
        Ela começa herdando da classe Gtk.ApplicationWindow
    '''
    def __init__(self, app):
        '''
            No metodo init tem 2 parâmetros o primeiro é o self e o segundo é app.
            Na estrutura de orientação a objetos do python é padrão que todos 
            os metodos de uma classe tem como primeiro parâmetro o self na declaração do metodo.
            Agora o parâmetro app é a instância app que chamou a janela.
        '''
        Gtk.Window.__init__(self, title="Gestão de compras",application=app)
        # Definindo o tamanho inicial da tela
        self.set_default_size(980,600)

        # Instânciando um Grid.O Grid ajudará a organizar os widget na tela
        self.grid = Gtk.Grid(column_homogeneous=True)
        self.add(self.grid)


        btn_produtos = Gtk.Button(label="Produto")            
        btn_compras = Gtk.Button(label="Compra")            
        btn_estoque = Gtk.Button(label="Estoque")         
        btn_estimativas = Gtk.Button(label="Estimativa")          

        btn_pesquisa = Gtk.Button(label="Pesquisar")
        btn_pesquisa.set_name("btnpesquisar")    

        lbl_produto = Gtk.Label("Produto")
        lbl_produto.set_name("title")  
        lbl_produto.set_halign(Gtk.Align.START)

        inp_nome_prod = Gtk.Entry()

        
        '''
            Algunas informações necessárias para entender o metodo attach do Grid.
            parametros:
            1º É um widget que eu quero posicionar na tela
            2º É o indice da coluna onde quero colocar o widget
            3º É o indice da linha onde quero colocar o widget
            4º É o numero de colunas que irá ocupar
            5º É o numero de linhas que irá ocupar
        '''

        # Esse segunda grid vai organizar o formulário
        self.grid2 = Gtk.Grid(column_homogeneous=True)

        self.grid2.attach(lbl_produto,0,0,4,1)
        self.grid2.attach(inp_nome_prod,0,1,3,1)
        self.grid2.attach(btn_pesquisa,3,1,1,1)
        #espaço entre as linhas do formulário
        self.grid2.set_row_spacing(10)
        #espaço entre as counas do formulário
        self.grid2.set_column_spacing(10)
        #margin do formulário
        self.grid2.props.margin = 20


        self.grid.attach(btn_produtos,0,0,1,1)
        self.grid.attach(btn_compras,1,0,1,1)
        self.grid.attach(btn_estoque,2,0,1,1)
        self.grid.attach(btn_estimativas,3,0,1,1)

        self.grid.attach(self.grid2,0,1,4,1)

        



class App(Gtk.Application):
    '''
       Classe responsável por gerenciar a applicação e o seu ciclo de vida.
       Apartir desta classe que iniciaremos a primeira janela.
    '''
    def __init__(self):
        Gtk.Application.__init__(self,
                                 application_id="app.home",
                                 flags=Gio.ApplicationFlags.FLAGS_NONE)

    def style(self):
        style_provider = Gtk.CssProvider()
        style_provider.load_from_path("style/app2.css")

        Gtk.StyleContext.add_provider_for_screen(
            Gdk.Screen.get_default(),
            style_provider,
            Gtk.STYLE_PROVIDER_PRIORITY_APPLICATION
        )        

   
    def do_activate(self):
        window = AppWindow(self)
        window.connect("delete-event", Gtk.main_quit)
        window.show_all()
        self.style()    
	

if __name__ == '__main__':
    app = App()
    app.run(sys.argv)