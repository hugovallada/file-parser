# File Parser

Projeto criado para gerar novos arquivos de textos, formatados a partir de um arquivo de texto.

## Args
    -fileNames -> String que contém os arquivos a serem lidos
        Os arquivos devem ser passados separados por "," ou passados dentro de "" com separação por espaço.
        Ex: -fileNames acesso.txt,acesso2.txt
            -fileNames "acesso.txt acesso2.txt"

    -newfileNames -> String que contém o novo nomes dos arquivos.
        Podem ser separados por "," ou passados dentro de "" com separação por espaço.
        Ex: -newFileNames acesso.txt,acesso2.txt
            -newFileNames "acesso.txt acesso2.txt"

    -parsers -> String com valores a serem substituidos e os valores substitutos.
        Podem ser separados por "," ou passados dentro de "" com separação por espaço.
        Devem ser escritos em chave/valor separados por =
        Ex: -parsers *=a,%=C
        Ex: -parsers "*=a %=C"

        Caso use caracteres com $, usar \$.

    -deleteOld -> Se usado, irá deletar os arquivos antigos
        
