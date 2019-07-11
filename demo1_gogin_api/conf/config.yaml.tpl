runmode: release                                # 开发模式, debug, release, test
addr: :8080                                     # HTTP绑定端口
name: apigo                                     # API Server的名字
checkurl: http://127.0.0.1:8080/check/health    # pingServer函数请求的API服务器的ip:port
max_ping_count: 10                              # pingServer函数try的次数
token_expired: 30                               # 登陆后token过期时间, 单位分钟
log:
    writers: file,stdout
    logger_level: DEBUG
    logger_file: log/apigo.log
    log_format_text: false
    rollingPolicy: size
    log_rotate_date: 1
    log_rotate_size: 1
    log_backup_count: 7
db:
    engine: mysql
    username: root
    password: password
    addr: 127.0.0.1:3306
    name: testdb
redis:
    addr: 127.0.0.1:6379
    password: 
    db: 0
    poolsize: 10

