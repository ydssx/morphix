CREATE DATABASE IF NOT EXISTS `order`;

USE `order`;

CREATE TABLE IF NOT EXISTS orders (
    order_id INT PRIMARY KEY COMMENT '订单ID',
    user_id INT COMMENT '客户ID',
    amount DECIMAL(10, 2) COMMENT '订单金额',
    status VARCHAR(50) COMMENT '订单状态',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS order_items (
    item_id INT PRIMARY KEY COMMENT '订单商品ID',
    order_id INT COMMENT '订单ID',
    product_id INT COMMENT '商品ID',
    quantity INT COMMENT '商品数量',
    price DECIMAL(10, 2) COMMENT '商品单价',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
    FOREIGN KEY (order_id) REFERENCES orders (order_id)
);

CREATE TABLE products (
    product_id INT PRIMARY KEY COMMENT '商品ID',
    name VARCHAR(100) COMMENT '商品名称',
    description VARCHAR(200) COMMENT '商品描述',
    price DECIMAL(10, 2) COMMENT '商品价格',
    status ENUM('AVAILABLE', 'UNAVAILABLE') COMMENT '商品状态 AVAILABLE:可用，UNAVAILABLE:不可用',
    image VARCHAR(200) COMMENT '商品图片',
    stock INT COMMENT '商品库存',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间'
);

INSERT INTO products (product_id, name, description, price, status, image, stock) VALUES
(1, '商品1', '商品1的描述', 10.00, 'AVAILABLE', 'https://example.com/image1.png', 100),
(2, '商品2', '商品2的描述', 20.00, 'AVAILABLE', 'https://example.com/image2.png', 200),
(3, '商品3', '商品3的描述', 30.00, 'UNAVAILABLE', 'https://example.com/image3.png', 0);
