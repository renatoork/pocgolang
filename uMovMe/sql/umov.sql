connect system;

CREATE USER MEGAUMOV IDENTIFIED BY MEGAUMOV;
GRANT "ROLE_MEGA" TO "MEGAUMOV";
ALTER USER "MEGAUMOV" DEFAULT ROLE ALL;

ALTER USER MEGAUMOV DEFAULT TABLESPACE TSD_MEGA
TEMPORARY TABLESPACE TEMP PROFILE DEFAULT QUOTA UNLIMITED ON 
TSD_MEGA;

GRANT ALL ON MGADM.EST_GRUPOS to MEGAUMOV;
GRANT ALL ON MGADM.EST_PRODUTOS to MEGAUMOV;
GRANT ALL ON MGGLO.GLO_AGENTES to MEGAUMOV;
GRANT ALL ON MGGLO.GLO_AGENTES_ID to MEGAUMOV;
grant all on MGVEN.VEN_PEDIDOVENDA to MEGAUMOV;
grant all on MGVEN.VEN_ITEMPEDIDOVENDA to MEGAUMOV;
grant all on MGVEN.VEN_TIPODOCUMENTO to MEGAUMOV;
grant all on MGWEB.VEN_PCK_PEDIDOVENDAWEB to MEGAUMOV
GRANT ALL ON MGGLO.glo_pck_agente TO MEGAUMOV;
grant all on mgven.Ven_Pck_Utilidades to megaumov;

connect MEGAUMOV;

/* GROUP */
drop table MEGAUMOV.UMOV_GROUP;
create table MEGAUMOV.UMOV_GROUP
(
  id                  number,
  operacao            char(1) default 'I',
  descricao           varchar2(400),
  codigoidentificador varchar2(20)
)
tablespace TSD_MEGA
  storage
  (
    initial 64K
    minextents 1
    maxextents unlimited
  )
nologging;
-- Add comments to the columns 
comment on column UMOV_GROUP.id
  is 'Id Sequencial';
comment on column UMOV_GROUP.operacao
  is 'Operação do registro (I-Inclusão; A-Alteração; E;Exclusão)';
comment on column UMOV_GROUP.descricao
  is 'Descrição do Grupo';
comment on column UMOV_GROUP.codigoidentificador
  is 'Código Identificador';
-- Create/Recreate primary, unique and foreign key constraints 
alter table UMOV_GROUP
  add constraint umov_pk_group primary key (ID);
-- Create/Recreate check constraints 
alter table UMOV_GROUP
  add constraint umov_Ck_groupoperacao
  check (operacao in ('I','A','E'));

alter table UMOV_GROUP
  add constraint UMOV_UK_GROUPDESC unique (DESCRICAO)
  using index ;  

  -- Create sequence 
create sequence MEGAUMOV.UMOV_SEQ_GROUP
minvalue 1
maxvalue 9999999999999999999999999999
start with 1
increment by 1
cache 20;

create or replace trigger MEGAUMOV.UMOV_TRA_GROUP
  before insert or update of GRU_ST_NOME or delete on MGADM.EST_GRUPOS
for each row
declare
  vOperacao char(1) := null;
  vDescricao varchar2(200);
  vCodIdentificador number;
begin

  if (INSERTING) then
    vOperacao := 'I';
    vDescricao := :new.GRU_ST_NOME;
    vCodIdentificador := :NEW.GRU_IN_CODIGO;
  end if;

  if (UPDATING) then
    if (:new.GRU_ST_NOME <> :old.GRU_ST_NOME) then
      vOperacao := 'A';
      vDescricao := :new.GRU_ST_NOME;
      vCodIdentificador := :NEW.GRU_IN_CODIGO;
    end if;
  end if;

  if (DELETING) then
    vOperacao := 'E';
    vDescricao := :old.GRU_ST_NOME;
    vCodIdentificador := :old.GRU_IN_CODIGO;
  end if;

  if (vOperacao is not null) then
    begin
      insert into MEGAUMOV.UMOV_GROUP (ID, OPERACAO, DESCRICAO, CODIGOIDENTIFICADOR)
                               values (MEGAUMOV.Umov_Seq_Group.Nextval, vOperacao, vDescricao, vCodIdentificador);
    exception
      when DUP_VAL_ON_INDEX then
        insert into MEGAUMOV.UMOV_GROUP (ID, OPERACAO, DESCRICAO, CODIGOIDENTIFICADOR)
                                 values (MEGAUMOV.Umov_Seq_Group.Nextval, vOperacao, vDescricao||' ('||TRIM(TO_CHAR(vCodIdentificador))||')', vCodIdentificador);
    end;
  end if;

end UMOV_TRA_GROUP;
/

/* ITEM */
drop table MEGAUMOV.UMOV_ITEM;
create table MEGAUMOV.UMOV_ITEM
(
  id                  number,
  operacao            char(1) default 'I',
  descricao           varchar2(400),
  codigoidentificador varchar2(20),
  especie             varchar2(3),
  peso                number(6,2),
  grupo               number(7)
)
tablespace TSD_MEGA
  storage
  (
    initial 64K
    minextents 1
    maxextents unlimited
  )
nologging;
-- Add comments to the columns 
comment on column UMOV_ITEM.id
  is 'Id Sequencial';
comment on column UMOV_ITEM.operacao
  is 'Operação do registro (I-Inclusão; A-Alteração; E;Exclusão)';
comment on column UMOV_ITEM.descricao
  is 'Descrição do Grupo';
comment on column UMOV_ITEM.codigoidentificador
  is 'Código Identificador';
comment on column UMOV_ITEM.especie
  is 'Espécie';
comment on column UMOV_ITEM.peso
  is 'Peso';
comment on column UMOV_ITEM.grupo
  is 'SubGrupo';
-- Create/Recreate primary, unique and foreign key constraints 
alter table UMOV_ITEM
  add constraint umov_pk_item primary key (ID);
-- Create/Recreate check constraints 
alter table UMOV_ITEM
  add constraint umov_Ck_itemoperacao
  check (operacao in ('I','A','E'));
alter table UMOV_ITEM
  add constraint UMOV_UK_ITEMDESC unique (DESCRICAO)
  using index ;    

  -- Create sequence 
create sequence MEGAUMOV.UMOV_SEQ_ITEM
minvalue 1
maxvalue 9999999999999999999999999999
start with 1
increment by 1
cache 20;

create or replace trigger UMOV_TRA_ITEM
  before insert or update of PRO_ST_DESCRICAO or delete on MGADM.EST_PRODUTOS
for each row
declare
  vOperacao char(1) := null;
  vDescricao varchar2(200);
  vCodIdentificador number;
  vEspecie varchar2(10);
  vPeso number(6,2);
  vGrupo number(7);
begin

  if (INSERTING) then
    vOperacao := 'I';
    vDescricao := :new.PRO_ST_DESCRICAO;
    vCodIdentificador := :new.PRO_IN_CODIGO;
    vEspecie := :new.UNIP_ST_UNIDADE;
    vPeso := :new.PRO_RE_PELIQUIDO;
    vGrupo := :new.GRU_IN_CODIGO;
  end if;

  if (UPDATING) then
    if (:new.PRO_ST_DESCRICAO <> :old.PRO_ST_DESCRICAO) then
      vOperacao := 'A';
      vDescricao := :new.PRO_ST_DESCRICAO;
      vCodIdentificador := :new.PRO_IN_CODIGO;
      vEspecie := :new.UNIP_ST_UNIDADE;
      vPeso := :new.PRO_RE_PELIQUIDO;
      vGrupo := :new.GRU_IN_CODIGO;
    end if;
  end if;

  if (DELETING) then
    vOperacao := 'E';
    vDescricao := :old.PRO_ST_DESCRICAO;
    vCodIdentificador := :old.PRO_IN_CODIGO;
    vEspecie := :old.UNIP_ST_UNIDADE;
    vPeso := :old.PRO_RE_PELIQUIDO;
    vGrupo := :old.GRU_IN_CODIGO;
  end if;

  if (vOperacao is not null) then
    begin
    insert into MEGAUMOV.UMOV_ITEM (ID, OPERACAO, DESCRICAO, CODIGOIDENTIFICADOR, ESPECIE, PESO, GRUPO)
                           values (MEGAUMOV.Umov_Seq_Item.Nextval, vOperacao, vDescricao, vCodIdentificador, vEspecie, vPeso, vGrupo);
    exception
      when DUP_VAL_ON_INDEX then
        insert into MEGAUMOV.UMOV_ITEM (ID, OPERACAO, DESCRICAO, CODIGOIDENTIFICADOR, ESPECIE, PESO, GRUPO)
                               values (MEGAUMOV.Umov_Seq_Item.Nextval, vOperacao, vDescricao||' ('||vCodIdentificador||')', vCodIdentificador, vEspecie, vPeso, vGrupo);
    end;
  end if;
end UMOV_TRA_ITEM;
/


/* CLIENTE */
drop table MEGAUMOV.UMOV_CLIENTE;
create table MEGAUMOV.UMOV_CLIENTE
(
  id                  number,
  operacao            char(1) default 'I',
  razaosocial         varchar2(200),
  nomefantasia        varchar2(200),
  codigoidentificador varchar2(20),
  cidade              varchar2(100),
  datacadastro        date
)
tablespace TSD_MEGA
  storage
  (
    initial 64K
    minextents 1
    maxextents unlimited
  )
nologging;
-- Add comments to the columns 
comment on column UMOV_CLIENTE.id
  is 'Id Sequencial';
comment on column UMOV_CLIENTE.operacao
  is 'Operação do registro (I-Inclusão; A-Alteração; E;Exclusão)';
comment on column UMOV_CLIENTE.razaosocial
  is 'Razão Social do Cliente';
comment on column UMOV_CLIENTE.nomefantasia
  is 'Nome Fantasia do Cliente';
comment on column UMOV_CLIENTE.codigoidentificador
  is 'Código Identificador';
comment on column UMOV_CLIENTE.cidade
  is 'Município do Cliente';
comment on column UMOV_CLIENTE.datacadastro
  is 'Data do Cadastro do Cliente';
-- Create/Recreate primary, unique and foreign key constraints 
alter table UMOV_CLIENTE
  add constraint umov_pk_cliente primary key (ID);
-- Create/Recreate check constraints 
alter table UMOV_CLIENTE
  add constraint umov_Ck_clienteoperacao
  check (operacao in ('I','A','E'));
alter table UMOV_CLIENTE
  add constraint UMOV_UK_CLIENTE unique (NOMEFANTASIA)
  using index ;    

  -- Create sequence 
create sequence MEGAUMOV.UMOV_SEQ_CLIENTE
minvalue 1
maxvalue 9999999999999999999999999999
start with 1
increment by 1
cache 20;

create or replace trigger UMOV_TRA_CLIENTE
  before insert or update of  AGN_ST_FANTASIA, AGN_ST_NOME, AGN_ST_MUNICIPIO, AGN_DT_ULTIMAATUCAD or delete on MGGLO.GLO_AGENTES
for each row
declare
  vOperacao char(1) := null;
  vCodIdentificador number;
  vrazaosocial         varchar2(200);
  vnomefantasia        varchar2(200);
  vcidade              varchar2(100);
  vdatacadastro        date;
begin

  if (INSERTING) then
    vOperacao := 'I';
    vCodIdentificador := :new.AGN_IN_CODIGO;
    vrazaosocial      := :new.AGN_ST_NOME;
    vnomefantasia     := :new.AGN_ST_FANTASIA;
    vcidade           := :new.AGN_ST_MUNICIPIO;
    vdatacadastro     := :new.AGN_DT_ULTIMAATUCAD;
  end if;

  if (UPDATING) then
    vOperacao := 'A';
    vCodIdentificador := :new.AGN_IN_CODIGO;
    vrazaosocial      := :new.AGN_ST_NOME;
    vnomefantasia     := :new.AGN_ST_FANTASIA;
    vcidade           := :new.AGN_ST_MUNICIPIO;
    vdatacadastro     := :new.AGN_DT_ULTIMAATUCAD;
  end if;

  if (DELETING) then
    vOperacao := 'E';
    vCodIdentificador := :old.AGN_IN_CODIGO;
    vrazaosocial      := :old.AGN_ST_NOME;
    vnomefantasia     := :old.AGN_ST_FANTASIA;
    vcidade           := :old.AGN_ST_MUNICIPIO;
    vdatacadastro     := :old.AGN_DT_ULTIMAATUCAD;
  end if;

  if (vOperacao is not null) then
    begin
    insert into MEGAUMOV.UMOV_CLIENTE (ID, OPERACAO, CODIGOIDENTIFICADOR, RAZAOSOCIAL, NOMEFANTASIA, CIDADE, DATACADASTRO)
                           values (MEGAUMOV.UMOV_SEQ_CLIENTE.Nextval, vOperacao, vCodIdentificador, vRazaoSocial, vNomeFantasia, vCidade, vDataCadastro);
    exception
      when DUP_VAL_ON_INDEX then
        insert into MEGAUMOV.UMOV_CLIENTE (ID, OPERACAO, CODIGOIDENTIFICADOR, RAZAOSOCIAL, NOMEFANTASIA, CIDADE, DATACADASTRO)
                               values (MEGAUMOV.UMOV_SEQ_CLIENTE.Nextval, vOperacao, vCodIdentificador, vRazaoSocial, vNomeFantasia||' ('||vCodIdentificador||')', vCidade, vDataCadastro);
    end;
  end if;
end UMOV_TRA_CLIENTE;
/

/* TipoDocumento */
drop table MEGAUMOV.UMOV_TIPODOCUMENTO;
create table MEGAUMOV.UMOV_TIPODOCUMENTO
(
  id                  number,
  operacao            char(1) default 'I',
  descricao           varchar2(400),
  codigoidentificador varchar2(20)
);
-- Add comments to the columns 
comment on column UMOV_TIPODOCUMENTO.id
  is 'Id Sequencial';
comment on column UMOV_TIPODOCUMENTO.operacao
  is 'Operação do registro (I-Inclusão; A-Alteração; E;Exclusão)';
comment on column UMOV_TIPODOCUMENTO.descricao
  is 'Descrição do Tipo de Documento';
comment on column UMOV_TIPODOCUMENTO.codigoidentificador
  is 'Código Identificador';
-- Create/Recreate primary, unique and foreign key constraints 
alter table UMOV_TIPODOCUMENTO
  add constraint umov_pk_TipoDoc primary key (ID);
-- Create/Recreate check constraints 
alter table UMOV_TIPODOCUMENTO
  add constraint umov_Ck_TipoDocOperacao
  check (operacao in ('I','A','E'));

alter table UMOV_TIPODOCUMENTO
  add constraint UMOV_UK_TIPODOCDESC unique (DESCRICAO)
  using index ;  

  -- Create sequence 
create sequence MEGAUMOV.UMOV_SEQ_TIPODOC
minvalue 1
maxvalue 9999999999999999999999999999
start with 1
increment by 1
cache 20;

create or replace trigger MEGAUMOV.UMOV_TRA_TIPODOC
  before insert or update of TPD_ST_DESCRICAO or delete on MGVEN.VEN_TIPODOCUMENTO
for each row
declare
  vOperacao char(1) := null;
  vDescricao varchar2(200);
  vCodIdentificador number;
begin

  if (INSERTING and :new.Tpd_Ch_Tipodocumento = 'P') then
    vOperacao := 'I';
    vDescricao := :new.TPD_ST_DESCRICAO;
    vCodIdentificador := :NEW.TPD_IN_CODIGO;
  end if;

  if (UPDATING and :new.Tpd_Ch_Tipodocumento = 'P') then
    if (:new.TPD_ST_DESCRICAO <> :old.TPD_ST_DESCRICAO) then
      vOperacao := 'A';
      vDescricao := :new.TPD_ST_DESCRICAO;
      vCodIdentificador := :NEW.TPD_IN_CODIGO;
    end if;
  end if;

  if (DELETING and :old.Tpd_Ch_Tipodocumento = 'P') then
    vOperacao := 'E';
    vDescricao := :old.TPD_ST_DESCRICAO;
    vCodIdentificador := :old.TPD_IN_CODIGO;
  end if;

  if (vOperacao is not null) then
    begin
      insert into MEGAUMOV.UMOV_TIPODOCUMENTO (ID, OPERACAO, DESCRICAO, CODIGOIDENTIFICADOR)
                               values (MEGAUMOV.UMOV_SEQ_TIPODOC.Nextval, vOperacao, vDescricao, vCodIdentificador);
    exception
      when DUP_VAL_ON_INDEX then
        insert into MEGAUMOV.UMOV_TIPODOCUMENTO (ID, OPERACAO, DESCRICAO, CODIGOIDENTIFICADOR)
                                 values (MEGAUMOV.UMOV_SEQ_TIPODOC.Nextval, vOperacao, vDescricao||' ('||TRIM(TO_CHAR(vCodIdentificador))||')', vCodIdentificador);
    end;
  end if;

end UMOV_TRA_TIPODOC;
/

/* Importa Cliente */
drop table MEGAUMOV.UMOV_IMPCLIENTE;
create table MEGAUMOV.UMOV_IMPCLIENTE
(
  id                  number,
  operacao            char(1) default 'I',
  razaosocial         varchar2(200),
  nomefantasia        varchar2(200),
  codigoidentificador varchar2(20),
  cidade              varchar2(100),
  datacadastro        date
);
-- Add comments to the columns 
comment on column UMOV_IMPCLIENTE.id
  is 'Id Sequencial';
comment on column UMOV_IMPCLIENTE.operacao
  is 'Operação do registro (I-Inclusão; A-Alteração; E;Exclusão)';
comment on column UMOV_IMPCLIENTE.razaosocial
  is 'Razão Social do Cliente';
comment on column UMOV_IMPCLIENTE.nomefantasia
  is 'Nome Fantasia do Cliente';
comment on column UMOV_IMPCLIENTE.codigoidentificador
  is 'Código Identificador';
comment on column UMOV_IMPCLIENTE.cidade
  is 'Município do Cliente';
comment on column UMOV_IMPCLIENTE.datacadastro
  is 'Data do Cadastro do Cliente';
alter table UMOV_IMPCLIENTE
  add constraint UMOV_UK_IMPCLIENTE unique (NOMEFANTASIA)
  using index ;    


create or replace package UMOV_PCK_INTEGRAUMOV is

  -- Author  : RENATO
  -- Created : 14/11/2014
  -- Purpose : Rotinas de integração com a plataforma uMOV.ME

  -- importa os clientes do uMOV para o Mega ERP.
  procedure p_ImportaAgente;
  
end UMOV_PCK_INTEGRAUMOV;
/
create or replace package body UMOV_PCK_INTEGRAUMOV is
  procedure p_ImportaAgente is
    vXml varchar2(32567);
  begin
    for i in (select * from MEGAUMOV.UMOV_IMPCLIENTE) loop
      vXml := '<Agente OPERACAO="I">';
      vXml := vXml || '<AGN_ST_NOME>'||i.RAZAOSOCIAL||'</AGN_ST_NOME>';
      vXml := vXml || '<AGN_ST_FANTASIA>'||i.NOMEFANTASIA||'</AGN_ST_FANTASIA>';
      vXml := vXml || '<AGN_ST_MUNICIPIO>'||i.CIDADE||'</AGN_ST_MUNICIPIO>';
      vXml := vXml || '<AgenteId>';
      vXml := vXml || '<AGN_TAU_ST_CODIGO>C</AGN_TAU_ST_CODIGO>';
      vXml := vXml || '</AgenteId>';
      vXml := vXml || '<Parametros>';
      vXml := vXml || '<FIL_IN_CODIGO>3</FIL_IN_CODIGO>';
      vXml := vXml || '</Parametros>';
      vXml := vXml || '</Agente>';
      DBMS_OUTPUT.PUT_LINE(vXML);
      MGGLO.glo_pck_agente.p_integra(vXML);
      delete MEGAUMOV.UMOV_IMPCLIENTE 
       where id = i.Id;
      commit;
    end loop;
  end;
end UMOV_PCK_INTEGRAUMOV;
/

create or replace trigger MEGAUMOV.VEN_TRA_ITEMPEDIDOVENDAPRECO
  before insert on ven_itempedidovenda  
  for each row
begin
  :new.Itp_Re_Valorunitario := 100;    
end VEN_TRA_ITEMPEDIDOVENDAPRECO;
/
