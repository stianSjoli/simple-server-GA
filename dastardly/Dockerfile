FROM public.ecr.aws/portswigger/dastardly:latest
WORKDIR /usr/local/burpsuite_enterprise/bin
ENV BURP_START_URL https://ginandjuice.shop/
ENV BURP_REPORT_FILE_PATH ../dastardly/dastardly-report.xml
RUN mkdir ../dastardly/
RUN touch ../dastardly/dastardly-report.xml 
RUN chmod go+w ../dastardly/dastardly-report.xml
ENTRYPOINT ./dastardly  
 
