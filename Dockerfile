FROM golang:1.22.3
LABEL authors="Osval"
# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app
# Copie os arquivos go.mod e go.sum para o contêiner
COPY go.mod .
COPY go.sum .

# Baixe as dependências do módulo Go
RUN go mod download

# Copie o restante do código fonte para o contêiner
COPY . .

# Compile a aplicação
RUN go build -o main .

# Defina a porta que será exposta
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./main"]
ENTRYPOINT ["top", "-b"]