find -type f -iname '*.sh' | sed -e 's/.sh//g' -e 's/[./]//g'
