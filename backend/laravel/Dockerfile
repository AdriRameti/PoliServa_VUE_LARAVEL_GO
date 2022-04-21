FROM php:8.0.14-fpm-alpine

RUN docker-php-ext-install pdo pdo_mysql sockets

RUN curl -sS https://getcompose.org/installer | php -- \ --install-dir=/usr/local/bin --filename=composer

COPY  --from=composer:latest /usr/bin/composer /usr/bin/composer

WORKDIR /app/laravel

COPY ./laravel .

COPY ./db/wait-for-it.sh .

RUN chmod +x wait-for-it.sh

RUN composer install
