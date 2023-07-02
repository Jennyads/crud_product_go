<h3>Anotações</h3>

O pacote "html/template" é usado para criar textos orientados a dados no HTML. Este pacote implementa templates controlados por dados para gerar saída em HTML segura contra injeção de código.

O pacote "net/http"da biblioteca padrão da Go, possibilita a criação de servidores. Esse pacote torna fácil a criação de servidores HTTP. Por exemplo o comando ListenAndServe() inicia um servidor HTTP pertence a este biblioteca.


O padrão MVC é uma forma de dividir a aplicação em três partes. Nesse padrão, dividimos as responsabilidades em três partes: modelo, visão e controlador.


Para recuperar o valor do id no controle de produtos, podemos usar idDoProduto := r.URL.Query().Get("id").

A rota HandleFunc("/update", controllers.Update) é responsável por atualizar um produto no banco de dados. Essa rota, é responsável por atualizar um produto no banco de dados com base nos valores alterados do edit.

A rota HandleFunc("/new", controllers.New) é responsável por abrir o formulário para criar um produto, porém a inserção no banco é feita pela rota HandleFunc("/insert", controllers.Insert). Na rota do new, abrimos o formulário. Já na rota do insert, buscamos os valores adicionados no new e pedimos para o modelo de produtos executar a inserção no banco.
