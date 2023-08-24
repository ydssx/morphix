CREATE TABLE quotes (
    id INT PRIMARY KEY AUTO_INCREMENT,  -- 报价ID（自增主键）
    product_id INT NOT NULL,           -- 产品ID
    price DECIMAL(10, 2) NOT NULL,      -- 报价价格
    date TIMESTAMP NOT NULL,           -- 报价日期
    customer_info VARCHAR(255),         -- 客户信息
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE discounts (
    id INT PRIMARY KEY AUTO_INCREMENT,      -- 折扣ID（自增主键）
    product_id INT NOT NULL,               -- 产品ID
    discount_percentage DECIMAL(5, 2) NOT NULL,  -- 折扣百分比
    CONSTRAINT fk_discount_product FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE coupons (
    id INT PRIMARY KEY AUTO_INCREMENT,    -- 优惠券ID（自增主键）
    code VARCHAR(50) NOT NULL,           -- 优惠券代码
    discount_type ENUM('percentage', 'amount') NOT NULL,  -- 折扣类型：百分比、固定金额
    discount_value DECIMAL(10, 2) NOT NULL,  -- 折扣值，根据折扣类型为百分比或金额
    expiration_date DATE NOT NULL         -- 优惠券过期日期
);

CREATE TABLE user_coupons (
    id INT PRIMARY KEY AUTO_INCREMENT,  -- 记录ID（自增主键）
    user_id INT NOT NULL,              -- 用户ID，关联到用户表
    coupon_id INT NOT NULL,            -- 优惠券ID，关联到优惠券表
    used BOOLEAN DEFAULT FALSE,        -- 是否已使用
    used_date TIMESTAMP,               -- 使用日期
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (coupon_id) REFERENCES coupons(id)
);
