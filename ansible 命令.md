## ��������
temp

	ansible yunce_search -m shell -a 'tail -n5 /var/log/spider/spider.log'

	ansible yunce_dp_alpha -m shell -a 'tail -n1 /var/log/data/data.log'

	ansible yunce_search -m shell -a 'rm -rf /usr/bin/chromedriver' -f20

	ansible yunce_search -m synchronize -a 'src=/root/packages/chromedriver dest=/usr/bin/' -f20

	ansible aos_spider -m shell -a 'chmod u+x,o+x /usr/bin/chromedriver' -f20
	

�鿴 Spider ��־
	
	# ���ز鿴
	ansible yunce_search -m shell -a 'tail -n5 /var/log/spider/spider.log' -f20

	# ʵʱ�鿴
	ansible yunce_search_alpha -m shell -a 'tail -20f /var/log/spider/spider.log'


DP ��־

�鿴 yunce_dp ��־

	# ���ز鿴
	ansible yunce_dp -m shell -a 'tail -n20 /var/log/dp/udp_receive.log'

	ansible yunce_dp -m shell -a 'tail -n20 /home/gpu/app/dp/alpha/GpuUdpClient/log/udp_log.log'

	# ʵʱ�鿴
	ansible yunce_dp -m shell -a 'tail -20f /home/gpu/app/dp/alpha/yunce_dp/log/data.log'

	ansible yunce_dp -m shell -a 'tail -20f /home/gpu/app/dp/alpha/GpuUdpClient/log/udp_log.log'
	

cd /home/gpu/app/dp/alpha/yunce_dp/log

Start

	ansible hosts -m shell -a 'supervisorctl start all' -f20

Stop

	ansible yunce_search_alpha -m shell -a 'supervisorctl stop all' -f10

Restart

	ansible yunce_search -m shell -a 'supervisorctl restart all' -f20

# ��������

	alpha:
	cd /root/projects/yunce/alpha/
	rm -rf yunce-search-spider/
	git clone -b alpha git@gitlab.web.zz:panziqiang/yunce-search-spider.git
	ansible-playbook /etc/ansible/yaml_dir/copy_projects_alpha.yaml -f10

	offical
	cd /root/projects/yunce/official
	rm -rf yunce-search-spider/
	git clone git@gitlab.web.zz:panziqiang/yunce-search-spider.git
	ansible-playbook /etc/ansible/yaml_dir/copy_projects_official.yaml -f20



### �Ʋ�Hosts
Spider
* ��ʽ yunce_search
* ���� yunce_search_alpha

DP
* ��ʽ 102��100
* ���� yunce_dp_alpha


ansible ����

ansible all -m shell -a 'systemctl enable supervisord'
ansible aos_dp -m shell -a 'docker stop $(docker ps -a -q)'

ansible yunce_dp_alpha -m shell -a 'supervisorctl start all'

ansible aos_dp -m shell -a 'supervisorctl status' -f10
ansible all -m shell -a 'mkdir /var/log/scrapy'

tail -n5 /var/log/data/data.log

ansible aos_dp -m shell -a 'tail -n10 /root/DataProcessing/log/data.log' -f10

ansible aos_dp -m synchronize -a 'src=/root/virtualenvs/dlqy_spider_packs.txt dest=/root/' -f10
ansible aos_task_root_queue -m synchronize -a 'src=/root/packages/chromedriver dest=/usr/bin/' -f20
ansible aos_dp -m shell -a 'mkdir /var/log/data' -f10


ansible yunce_search_alpha -m synchronize -a 'src=/root/projects/server_conf/yunce_alpha/supervisord.conf dest=/etc/'
ansible yunce_dp_alpha -m shell -a 'supervisorctl stop all'

ansible all -m shell -a 'supervisorctl '

ansible yunce_search_alpha -m shell -a 'rm -rf /var/log/spider*'



ansible yunce_search_alpha -m shell -a 'tail -n5 /var/log/spider/spider.log' 


ʵʱ�鿴 supervisor ��־ 20 ��
tail -20f /var/log/spider/spider.log

ʵʱ�鿴 scrapy ��־ 20��
tail -20f /var/log/spider-2019-8-12.log


tail -1f /var/log/data/data.log �� ʵʱ�鿴

�鿴 scrapy ��־
	
	ansible yunce_search -m shell -a 'tail -n20 /var/log/spider*.log' -f 50


	ansible yunce_search_alpha -m shell -a 'tail -n50 /var/log/spider*.log' -f 10



# �鿴 ���100�е�ǰ65��

	ansible yunce_search -m shell -a 'tail -n 100 /var/log/spider*.log | head -n 65' -f 10


��������
ֹͣ docker ansible yunce_search -m shell -a 'docker stop $(docker ps -a -q)' -f 50


���� logrotate ��־
ansible yunce_search -m shell -a 'logrotate /etc/logrotate.conf' -f 50


״̬�鿴
ansible yunce_search -m shell -a 'supervisorctl restart all'

ansible yunce_timing -m shell -a 'systemctl enable supervisord'



�ϴ������ļ�
ansible yunce_search -m copy -a 'src=/Users/jason/Downloads/sina_spider_m.py dest=/root/DLQYSpider2/DLQYSpider2/other_spiders'