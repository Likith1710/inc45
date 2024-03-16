FROM wordpress:latest

# Install Git and PHP_CodeSniffer
RUN apt-get update && \
    apt-get install -y git && \
    git clone https://github.com/squizlabs/PHP_CodeSniffer.git /usr/local/phpcs && \
    ln -s /usr/local/phpcs/bin/phpcs /usr/local/bin/phpcs

