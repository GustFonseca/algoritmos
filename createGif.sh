#!/bin/bash

# Verifica se o diretório existe e contém arquivos PNG
if [ ! -d "graphics" ] || [ -z "$(ls graphics/iter_*.png 2>/dev/null)" ]; then
    echo "Erro: O diretório 'graphics' não contém imagens PNG compatíveis."
    exit 1
fi

# Conta o número de imagens
count=$(ls graphics/iter_*.png | wc -l)

# Define uma largura mínima e uma largura máxima
min_width=800
max_width=1600

# Ajusta a largura com base na contagem, garantindo que não fique abaixo do mínimo
size=$((max_width / count))
if [ $size -lt $min_width ]; then
    size=$min_width
fi

echo "Criando GIF com largura ajustada para $size pixels."

# Pega a última imagem
last_image=$(ls graphics/iter_*.png | tail -n 1)

# Cria o GIF com atraso regular para todas as imagens, exceto a última, que terá um delay de 5 segundos
convert -resize ${size}x -delay 10 graphics/iter_*.png animation.gif

# Verifica se o GIF foi criado com sucesso
if [ $? -eq 0 ]; then
    echo "GIF criado com sucesso: animation.gif"
      rm graphics/iter_*.png
else
    echo "Erro ao criar o GIF."
fi
