FROM flyway/flyway

# Copie seus scripts de migração
COPY ./migration /flyway/sql
ENV FLYWAY_LOCATIONS=filesystem:/flyway/sql

ENTRYPOINT ["flyway", "migrate"]