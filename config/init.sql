
-- 应用信息
CREATE TABLE IF NOT EXISTS app(
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  app_id VARCHAR(255) UNIQUE COMMENT '应用ID',
  app_key VARCHAR(255) COMMENT '应用KEY',
  app_name VARCHAR(255) COMMENT '应用名称',
  app_desc VARCHAR(1000) COMMENT '应用描述',
  create_time TIMESTAMP COMMENT '创建时间',
  update_time TIMESTAMP COMMENT '更新时间',
  status int COMMENT '应用状态 0.待审核 1.已审核'

) CHARACTER SET utf8;

-- 用户表
CREATE TABLE IF NOT EXISTS users(
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  r_id BIGINT COMMENT '关联ID',
  app_id BIGINT COMMENT '应用ID',
  open_id VARCHAR(255) COMMENT '用户OPENID',
  create_time TIMESTAMP COMMENT '创建时间',
  status INT COMMENT '用户状态'
) CHARACTER SET utf8;