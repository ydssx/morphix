CREATE TABLE orders (
    order_id INT PRIMARY KEY COMMENT '订单ID',
    user_id INT COMMENT '客户ID',
    amount DECIMAL(10, 2) COMMENT '订单金额',
    status VARCHAR(50) COMMENT '订单状态'
);

CREATE TABLE order_items (
    item_id INT PRIMARY KEY COMMENT '订单商品ID',
    order_id INT COMMENT '订单ID',
    product_id INT COMMENT '商品ID',
    quantity INT COMMENT '商品数量',
    price DECIMAL(10, 2) COMMENT '商品单价',
    FOREIGN KEY (order_id) REFERENCES orders (order_id)
);

CREATE TABLE products (
    product_id INT PRIMARY KEY COMMENT '商品ID',
    name VARCHAR(100) COMMENT '商品名称',
    description VARCHAR(200) COMMENT '商品描述',
    price DECIMAL(10, 2) COMMENT '商品价格'
);
