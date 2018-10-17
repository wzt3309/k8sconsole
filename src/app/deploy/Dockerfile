FROM scratch
LABEL author=wzt3309
LABEL email=wzt3309@gmail.com

# Add All files from current working directory to the root of the image,
# e.g., copy dist dir
ADD . /

# The port the application default listen to
EXPOSE 8080
ENTRYPOINT ["./console"]