



iptables -t nat -A POSTROUTING -s 10.1.1.1/26 =j SNAT ==to-source 124.56.12.75




iptables -A input -p tcp --dport 80 -j ACCEPT

iptables -A input -p tcp -j DROP

iptables -A input -s 192.168.10.10 -p tcp -m multiport --dport 22,23,80,443 -j DROP



