#
# Generate the certificates and keys for testing.
#


# Generate the openssl configuration files.
cat > ca_cert.conf << EOF
[ req ]
default_bits       = 4096
default_md         = sha512
prompt             = no
encrypt_key        = no
distinguished_name = req_distinguished_name

[ req_distinguished_name ]
organizationName       = "My Company"             # O=
EOF

cat > server_cert.conf << EOF
[ req ]
default_bits       = 4096
default_md         = sha512
prompt             = no
encrypt_key        = no
distinguished_name = req_distinguished_name

[ req_distinguished_name ]
countryName            = "DE"                     # C=
localityName           = "Berlin"                 # L=
organizationName       = "My Company"             # O=
organizationalUnitName = "Departement"            # OU=
commonName             = "localhost"              # CN=
emailAddress           = "server@ca.com"          # CN/emailAddress=
EOF

cat > client_cert.conf << EOF
[ req ]
default_bits       = 4096
default_md         = sha512
prompt             = no
encrypt_key        = no
distinguished_name = req_distinguished_name

[ req_distinguished_name ]
countryName            = "DE"                     # C=
localityName           = "Berlin"                 # L=
organizationName       = "My Company"             # O=
organizationalUnitName = "Departement"            # OU=
commonName             = "localhost"              # CN=
emailAddress           = "client@ca.com"          # CN/emailAddress=
EOF

mkdir ca
mkdir server
mkdir client
mkdir certDER

# private key generation
openssl genrsa -out ca.key 2048
openssl genrsa -out server.key 2048
openssl genrsa -out client.key 2048

# cert requests
openssl req -out ca.req -key ca.key -new \
            -config ./ca_cert.conf
openssl req -out server.req -key server.key -new \
            -config ./server_cert.conf
openssl req -out client.req -key client.key -new \
            -config ./client_cert.conf

# generate the actual certs
openssl x509 -req -in ca.req -out ca.crt \
            -days 5000 -signkey ca.key
openssl x509 -req -in server.req -out server.crt \
            -CAcreateserial -days 5000 \
            -CA ca.crt -CAkey ca.key
openssl x509 -req -in client.req -out client.crt \
            -CAcreateserial -days 5000 \
            -CA ca.crt -CAkey ca.key

openssl x509 -in ca.crt -outform DER -out ca.der
openssl x509 -in server.crt -outform DER -out server.der
openssl x509 -in client.crt -outform DER -out client.der

mv ca.crt ca.key ca/
mv server.crt server.key server/
mv client.crt client.key client/
mv ca.der server.der client.der certDER/

rm *.conf
rm *.req
rm *.srl

# verify certs
openssl x509 -text -noout -in client/client.crt
openssl x509 -text -noout -in server/server.crt

openssl verify -verbose -CAfile ca/ca.crt client/client.crt
openssl verify -verbose -CAfile ca/ca.crt server/server.crt
