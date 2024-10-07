#!/bin/bash

# Conta o número de imagens
count=$(ls graphics/quicksort_iter_*.png | wc -l)

# Define uma largura mínima e uma largura máxima
min_width=800
max_width=1600

# Ajusta a largura com base na contagem, garantindo que não fique abaixo do mínimo
size=$((max_width / count))
if [ $size -lt $min_width ]; then
    size=$min_width
fi

# Cria o GIF
convert -delay 10 graphics/quicksort_iter_*.png quicksort_animation.gif
