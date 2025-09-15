CREATE TABLE meeting_entities (
    meeting_id     VARCHAR(64)  NOT NULL COMMENT '会议ID',
    meeting_name   VARCHAR(255) NOT NULL COMMENT '会议名称',
    meeting_password VARCHAR(64) DEFAULT NULL COMMENT '会议密码',
    description    TEXT         COMMENT '会议描述',
    host_id        VARCHAR(64)  NOT NULL COMMENT '主持人ID',
    host_name      VARCHAR(64)  NOT NULL COMMENT '主持人名字',
    start_time     BIGINT       NOT NULL COMMENT '开始时间，时间戳',
    end_time       BIGINT       COMMENT '结束时间，时间戳',
    PRIMARY KEY (meeting_id),
    KEY idx_meeting_id (meeting_id)
)COMMENT='会议表';

CREATE TABLE meeting_history (
  meeting_id VARCHAR(64) NOT NULL COMMENT '会议id',
  user_id VARCHAR(50) NOT NULL COMMENT '用户id',
  PRIMARY KEY (meeting_id, user_id),                        -- 联合主键
  CONSTRAINT fk_meeting FOREIGN KEY (meeting_id) 
    REFERENCES meeting_entities(meeting_id) 
    ON DELETE CASCADE ON UPDATE CASCADE
);
