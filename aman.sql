CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role ENUM('helper', 'pengguna') DEFAULT 'helper',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    SKU VARCHAR(10),
    size VARCHAR(10),
    color VARCHAR(50),
    image VARCHAR(225),
    price FLOAT NOT NULL,
    stock INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TABLE orders (
    id VARCHAR(100) PRIMARY KEY,
    user_id INT NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    customer_name VARCHAR(100) NOT NULL,
    customer_phone VARCHAR(15),
    customer_address TEXT, 
    status ENUM('pending', 'done') DEFAULT 'pending',
    expired_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE order_items (
    id VARCHAR(100) PRIMARY key NOT NULL,
    order_id VARCHAR(100) NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);


INSERT INTO products (name, SKU, size, color, image, price, stock) VALUES
('T-Shirt Casual', 'TS001', 'M', 'Merah', 'images/tshirt_casual_red.jpg', 199000, 150),
('T-Shirt Graphic', 'TS002', 'L', 'Biru', 'images/tshirt_graphic_blue.jpg', 249000, 100),
('Jeans Slim Fit', 'JN001', '32', 'Hitam', 'images/jeans_slimfit_black.jpg', 399000, 80),
('Jeans Regular Fit', 'JN002', '34', 'Indigo', 'images/jeans_regularfit_indigo.jpg', 349000, 90),
('Hoodie', 'HD001', 'L', 'Abu-abu', 'images/hoodie_gray.jpg', 499000, 70),
('Sweater Knit', 'SW001', 'M', 'Hijau', 'images/sweater_knit_green.jpg', 399000, 60),
('Jacket Denim', 'JK001', 'M', 'Biru', 'images/jacket_denim_blue.jpg', 699000, 50),
('Jacket Bomber', 'JK002', 'L', 'Hitam', 'images/jacket_bomber_black.jpg', 799000, 40),
('Shorts Chino', 'ST001', 'M', 'Coklat', 'images/shorts_chino_brown.jpg', 299000, 100),
('Skirt A-Line', 'SK001', 'S', 'Merah', 'images/skirt_a_line_red.jpg', 249000, 80),
('Dress Maxi', 'DR001', 'M', 'Biru', 'images/dress_maxi_blue.jpg', 599000, 30),
('Dress Midi', 'DR002', 'L', 'Putih', 'images/dress_midi_white.jpg', 499000, 40),
('Polo Shirt', 'PS001', 'M', 'Hitam', 'images/polo_shirt_black.jpg', 299000, 120),
('Tank Top', 'TT001', 'S', 'Merah', 'images/tank_top_red.jpg', 149000, 150),
('Cardigan', 'CD001', 'M', 'Krem', 'images/cardigan_cream.jpg', 399000, 60),
('Blazer Formal', 'BL001', 'L', 'Hitam', 'images/blazer_formal_black.jpg', 899000, 25),
('Belt Kulit', 'BL002', 'M', 'Coklat', 'images/belt_leather_brown.jpg', 199000, 90),
('Socks Panjang', 'SK002', 'M', 'Abu-abu', 'images/socks_long_gray.jpg', 50000, 200),
('Sandal Casual', 'SD001', '42', 'Hitam', 'images/sandal_casual_black.jpg', 199000, 150),
('Sneakers Sport', 'SN001', '42', 'Putih', 'images/sneakers_sport_white.jpg', 599000, 80),
('Boots Kulit', 'BT001', '41', 'Coklat', 'images/boots_leather_brown.jpg', 799000, 40),
('Kemeja Lengan Panjang', 'KM001', 'M', 'Biru', 'images/shirt_long_sleeve_blue.jpg', 349000, 70),
('Kemeja Lengan Pendek', 'KM002', 'L', 'Putih', 'images/shirt_short_sleeve_white.jpg', 299000, 90),
('Jumpsuit', 'JS001', 'M', 'Hitam', 'images/jumpsuit_black.jpg', 499000, 30),
('Romper', 'RP001', 'S', 'Merah', 'images/romper_red.jpg', 399000, 50),
('Kaos Kaki', 'SK003', 'M', 'Biru', 'images/socks_blue.jpg', 30000, 200),
('Kain Pashmina', 'KP001', 'One Size', 'Coklat', 'images/pashmina_brown.jpg', 150000, 100),
('Kain Scarf', 'SC002', 'One Size', 'Hijau', 'images/scarf_green.jpg', 120000, 80),
('Kemeja Flanel', 'KM003', 'M', 'Merah', 'images/fl anel_shirt_red.jpg', 399000, 60),
('Kemeja Batik', 'KM004', 'L', 'Coklat', 'images/batik_shirt_brown.jpg', 499000, 40),
('Sweater Hoodie', 'SW002', 'M', 'Hitam', 'images/sweater_hoodie_black.jpg', 599000, 30),
('Jacket Parka', 'JK003', 'L', 'Hijau', 'images/jacket_parka_green.jpg', 899000, 20),
('Celana Pendek', 'CP001', 'M', 'Biru', 'images/shorts_blue.jpg', 249000, 100),
('Celana Panjang', 'CP002', 'L', 'Hitam', 'images/pants_black.jpg', 399000, 80),
('Kemeja Casual', 'KM005', 'M', 'Putih', 'images/casual_shirt_white.jpg', 299000, 90),
('Dress Casual', 'DR003', 'M', 'Biru', 'images/casual_dress_blue.jpg', 499000, 50),
('Kain Syal', 'KS001', 'One Size', 'Merah', 'images/scarf_red.jpg', 150000, 100),
('Kain Selendang', 'KS002', 'One Size', 'Kuning', 'images/selendang_yellow.jpg', 120000, 80),
('Kemeja Oversized', 'KM006', 'L', 'Abu-abu', 'images/oversized_shirt_gray.jpg', 399000, 60),
('Celana Jogger', 'CJ001', 'M', 'Hitam', 'images/jogger_pants_black.jpg', 499000, 40),
('Kemeja Lengan Pendek Casual', 'KM007', 'M', 'Biru', 'images/casual_short_sleeve_blue.jpg', 299000, 90),
('Dress Mini', 'DR004', 'S', 'Merah', 'images/mini_dress_red.jpg', 399000, 50),
('Kemeja Formal', 'KM008', 'L', 'Putih', 'images/formal_shirt_white.jpg', 499000, 30);


