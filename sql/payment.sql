-- 订单表，存储订单信息
CREATE TABLE orders (
  id INT PRIMARY KEY,  -- 订单ID
  user_id INT NOT NULL,  -- 用户ID，关联用户表
  amount DECIMAL(10, 2) NOT NULL,  -- 订单金额
  currency VARCHAR(3) NOT NULL,  -- 货币代码
  status VARCHAR(20) NOT NULL,  -- 订单状态
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 创建时间
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- 更新时间
  FOREIGN KEY (user_id) REFERENCES users(id),  -- 外键关联用户表
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

-- 支付表，存储支付信息
CREATE TABLE payments (
  id INT PRIMARY KEY,  -- 支付ID
  order_id INT NOT NULL,  -- 订单ID，关联订单表
  amount DECIMAL(10, 2) NOT NULL,  -- 支付金额
  currency VARCHAR(3) NOT NULL,  -- 货币代码
  status VARCHAR(20) NOT NULL,  -- 支付状态
  transaction_id VARCHAR(50) NOT NULL,  -- 交易ID
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 创建时间
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- 更新时间
  FOREIGN KEY (order_id) REFERENCES orders(id)  -- 外键关联订单表
);

-- 退款表，存储退款信息
CREATE TABLE refunds (
  id INT PRIMARY KEY,  -- 退款ID
  payment_id INT NOT NULL,  -- 支付ID，关联支付表
  amount DECIMAL(10, 2) NOT NULL,  -- 退款金额
  currency VARCHAR(3) NOT NULL,  -- 货币代码
  status VARCHAR(20) NOT NULL,  -- 退款状态
  refund_id VARCHAR(50) NOT NULL,  -- 退款ID
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 创建时间
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- 更新时间
  FOREIGN KEY (payment_id) REFERENCES payments(id)  -- 外键关联支付表
);

