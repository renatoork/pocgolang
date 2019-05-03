-- exporta uMov.Me - GRUPO
WITH VENDAS AS (
SELECT DISTINCT 
       i.pro_pad_in_codigo, i.pro_in_codigo, p.Pro_St_Descricao, 
       p.Gru_Pad_In_Codigo, p.Gru_Ide_St_Codigo, p.Gru_In_Codigo, g.Gru_St_Nome, g.Gru_St_Defitem, t.Def_St_Descricao
  FROM mgven.ven_itempedidovenda i,
       mgadm.est_produtos p,
       mgadm.est_grupos g,
       MGADM.EST_DEFINICAOITEM t
 WHERE p.Pro_Tab_In_Codigo = i.Pro_Tab_In_Codigo
   AND p.Pro_pad_In_Codigo = i.Pro_pad_In_Codigo
   AND p.Pro_In_Codigo = i.Pro_In_Codigo
   AND g.gru_Tab_In_Codigo = p.gru_Tab_In_Codigo
   AND g.gru_pad_In_Codigo = p.gru_pad_In_Codigo
   AND g.gru_ide_st_Codigo = p.gru_ide_st_Codigo
   AND g.gru_In_Codigo = p.gru_In_Codigo
   AND t.def_ch_codigo = g.Gru_St_Defitem
)

SELECT XMLElement("group", XMLForest('true' active, Def_St_Descricao description, Gru_st_defitem alternativeidentifier)) resultado
/*SELECT DISTINCT 'true' active,
                Def_St_Descricao description,
                Gru_st_defitem alternativeidentifier
*/  FROM VENDAS

<group><active>true</active><description>componente</description><alternativeidentifier>co</alternativeidentifier></group>
