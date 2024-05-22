mongodump:                      MongoDB的Dump脚本
mysqldump:                      mysql的Dump脚本
restic:                         restic工具
restic_backup_tool:             备份工具
restic_backup_tool.yaml:        备份工具配置文件
restic_repo_password:           restic存储库密码

# 启用restic_backup_tool备份工具步骤如下：
#================1、配置环境变量==================
vim /etc/profile
...
export AWS_ACCESS_KEY_ID="minio"
export AWS_SECRET_ACCESS_KEY="minio123"
export RESTIC_REPOSITORY="s3:http://192.168.2.249:9000/restic"
...
source /etc/profile
# 验证环境变量是否生效
echo $RESTIC_REPOSITORY

#================2、初始化存储库==================
# 当配置了$RESTIC_REPOSITORY环境变量
restic init
# 当未配置$RESTIC_REPOSITORY环境变量
restic -r s3:http://192.168.2.249:9000/restic init

#================3、配置存储库密码=================
# 将密码保存在文件中
echo '******' > restic_repo_password
# 或者将密码配置为环境变量
vim /etc/profile
...
export RESTIC_PASSWORD="*******"

#================4、备份工具配置文件================
vim restic_backup_tool.yaml

#================5、启用备份工具===================
./restic_backup_tool
