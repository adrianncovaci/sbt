nmcli radio all | /usr/bin/awk '{print} END {if (NR == 0) print none}'