
FROM php:8.3-fpm AS builder

WORKDIR /var/www

RUN apt-get update && apt-get install -y \
    libpng-dev \
    libjpeg-dev \
    libfreetype6-dev \
    zip unzip git curl \
    libonig-dev libxml2-dev \
    && docker-php-ext-install pdo mbstring gd pdo_mysql mysqli pcntl

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

RUN chown -R www-data:www-data /var/www

# ===== Final =====
FROM php:8.3-fpm

WORKDIR /var/www

RUN apt-get update && apt-get install -y \
    libpng-dev \
    libjpeg-dev \
    libfreetype6-dev \
    libonig-dev \
    libxml2-dev \
    && docker-php-ext-install pdo mbstring gd pdo_mysql mysqli pcntl

COPY --from=builder /var/www /var/www
RUN chown -R www-data:www-data /var/www

EXPOSE 9000
CMD ["php-fpm"]