## 常用命令
temp

	ansible yunce_search -m shell -a 'tail -n5 /var/log/spider/spider.log'

	ansible yunce_dp_alpha -m shell -a 'tail -n1 /var/log/data/data.log'

	ansible yunce_search -m shell -a 'rm -rf /usr/bin/chromedriver' -f20

	ansible yunce_search -m synchronize -a 'src=/root/packages/chromedriver dest=/usr/bin/' -f20

	ansible aos_spider -m shell -a 'chmod u+x,o+x /usr/bin/chromedriver' -f20
	

查看 Spider 日志
	
	# 本地查看
	ansible yunce_search -m shell -a 'tail -n5 /var/log/spider/spider.log' -f20

	# 实时查看
	ansible yunce_search_alpha -m shell -a 'tail -20f /var/log/spider/spider.log'


DP 日志

查看 yunce_dp 日志

	# 本地查看
	ansible yunce_dp -m shell -a 'tail -n20 /var/log/dp/udp_receive.log'

	ansible yunce_dp -m shell -a 'tail -n20 /home/gpu/app/dp/alpha/GpuUdpClient/log/udp_log.log'

	# 实时查看
	ansible yunce_dp -m shell -a 'tail -20f /home/gpu/app/dp/alpha/yunce_dp/log/data.log'

	ansible yunce_dp -m shell -a 'tail -20f /home/gpu/app/dp/alpha/GpuUdpClient/log/udp_log.log'
	

cd /home/gpu/app/dp/alpha/yunce_dp/log

Start

	ansible hosts -m shell -a 'supervisorctl start all' -f20

Stop

	ansible yunce_search_alpha -m shell -a 'supervisorctl stop all' -f10

Restart

	ansible yunce_search -m shell -a 'supervisorctl restart all' -f20

# 更新爬虫

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



### 云测Hosts
Spider
* 正式 yunce_search
* 测试 yunce_search_alpha

DP
* 正式 102，100
* 测试 yunce_dp_alpha


ansible 命令

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


实时查看 supervisor 日志 20 行
tail -20f /var/log/spider/spider.log

实时查看 scrapy 日志 20行
tail -20f /var/log/spider-2019-8-12.log


tail -1f /var/log/data/data.log ＃ 实时查看

查看 scrapy 日志
	
	ansible yunce_search -m shell -a 'tail -n20 /var/log/spider*.log' -f 50


	ansible yunce_search_alpha -m shell -a 'tail -n50 /var/log/spider*.log' -f 10



# 查看 最后100行的前65行

	ansible yunce_search -m shell -a 'tail -n 100 /var/log/spider*.log | head -n 65' -f 10


常用命令
停止 docker ansible yunce_search -m shell -a 'docker stop $(docker ps -a -q)' -f 50


开启 logrotate 日志
ansible yunce_search -m shell -a 'logrotate /etc/logrotate.conf' -f 50


状态查看
ansible yunce_search -m shell -a 'supervisorctl restart all'

ansible yunce_timing -m shell -a 'systemctl enable supervisord'



上传单个文件
ansible yunce_search -m copy -a 'src=/Users/jason/Downloads/sina_spider_m.py dest=/root/DLQYSpider2/DLQYSpider2/other_spiders'