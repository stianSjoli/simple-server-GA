FROM public.ecr.aws/portswigger/dastardly:latest
WORKDIR /usr/local/burpsuite_enterprise/bin
ENV BURP_START_URL https://ginandjuice.shop/
ENV BURP_REPORT_FILE_PATH ../dastardly2/dastardly-report.xml
RUN mkdir ../dastardly2/
RUN touch ../dastardly2/dastardly-report.xml 
RUN chmod go+w ../dastardly2/dastardly-report.xml
ENTRYPOINT ./dastardly  
 
