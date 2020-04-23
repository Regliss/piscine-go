find -type f -iname '*.sh' | sed -e 's/.sh//g' | rev | cut -d'/' -f1 | rev
