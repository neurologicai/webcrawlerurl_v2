package site_map_test

import (
	"log"
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/site_map"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestSiteMap_Call(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.teste.com", httpmock.NewStringResponder(200, `
	<?xml version="1.0" encoding="UTF-8"?>
	<urlset
				xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
				xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
				xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9
							http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
		<url>
			<loc>https://leadster.com.br/</loc>
			<lastmod>2022-11-08T18:29:20+00:00</lastmod>
			<priority>1.00</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/</loc>
			<lastmod>2022-11-08T18:29:21+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-orcamentos-b2b/</loc>
			<lastmod>2022-11-08T18:29:22+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/formularios-vs-leadster/</loc>
			<lastmod>2022-11-08T18:29:23+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/botao-do-whatsapp-vs-leadster/</loc>
			<lastmod>2022-11-08T18:29:23+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chat-online-vs-leadster/</loc>
			<lastmod>2022-11-08T18:29:24+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-vs-leadster/</loc>
			<lastmod>2022-11-08T18:29:25+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-leads/agencias/</loc>
			<lastmod>2022-11-08T18:29:26+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-leads/b2b/</loc>
			<lastmod>2022-11-08T18:29:26+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-leads/educacional/</loc>
			<lastmod>2022-11-08T18:29:26+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-leads/imobiliaria/</loc>
			<lastmod>2022-11-08T18:29:27+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-leads/investimento/</loc>
			<lastmod>2022-11-08T18:29:28+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-leads/juridico/</loc>
			<lastmod>2022-11-08T18:29:28+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-leads/software/</loc>
			<lastmod>2022-11-08T18:29:29+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-leads/cases/</loc>
			<lastmod>2022-11-08T18:29:30+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/preco/</loc>
			<lastmod>2022-11-08T18:29:30+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/materiais/</loc>
			<lastmod>2022-11-08T18:29:32+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/marketing-conversacional/</loc>
			<lastmod>2022-11-08T18:29:32+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/panorama-de-geracao-de-leads-no-brasil/</loc>
			<lastmod>2022-11-08T18:29:32+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/midia-paga/</loc>
			<lastmod>2022-11-08T18:29:33+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/agencias-parceiras/programa/</loc>
			<lastmod>2022-11-08T18:29:34+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/agencias-parceiras/</loc>
			<lastmod>2022-11-08T18:29:34+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/como-funciona/</loc>
			<lastmod>2022-11-08T18:29:35+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/contato/</loc>
			<lastmod>2022-11-08T18:29:36+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/botao-do-whatsapp/</loc>
			<lastmod>2022-11-08T18:29:36+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/termos_de_uso/</loc>
			<lastmod>2022-11-08T18:29:36+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/privacidade/</loc>
			<lastmod>2022-11-08T18:29:37+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/politica-de-cookies/</loc>
			<lastmod>2022-11-08T18:29:37+00:00</lastmod>
			<priority>0.80</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/recurso-para-multas/</loc>
			<lastmod>2022-11-08T18:29:37+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/log10/</loc>
			<lastmod>2022-11-08T18:29:38+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/rafael-cassio/</loc>
			<lastmod>2022-11-08T18:29:38+00:00</lastmod>
			<priority>0.64</priority>
		</url>

		<url>
			<loc>https://leadster.com.br/?utm_source=blog&amp;utm_medium=menu&amp;utm_campaign=logo</loc>
			<lastmod>2022-11-08T18:29:41+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=menu&amp;utm_campaign=conhecer_leadster</loc>
			<lastmod>2022-11-08T18:29:42+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/materiais/plano-de-midia-paga/</loc>
			<lastmod>2022-11-08T18:30:19+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/materiais/framework-de-publicos/</loc>
			<lastmod>2022-11-08T18:30:20+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-leads/marketing-conversacional/</loc>
			<lastmod>2022-11-08T18:30:20+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/panorama-de-geracao-de-leads-no-brasil/b2b/</loc>
			<lastmod>2022-11-08T18:30:21+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/panorama-de-geracao-de-leads-no-brasil/concessionarias/</loc>
			<lastmod>2022-11-08T18:30:21+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/panorama-de-geracao-de-leads-no-brasil/educacional/</loc>
			<lastmod>2022-11-08T18:30:21+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/panorama-de-geracao-de-leads-no-brasil/imobiliarias/</loc>
			<lastmod>2022-11-08T18:30:22+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/panorama-de-geracao-de-leads-no-brasil/servicos/</loc>
			<lastmod>2022-11-08T18:30:22+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/panorama-de-geracao-de-leads-no-brasil/software/</loc>
			<lastmod>2022-11-08T18:30:23+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/panorama-de-geracao-de-leads-no-brasil/venda-de-produtos/</loc>
			<lastmod>2022-11-08T18:30:23+00:00</lastmod>
			<priority>0.64</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-leads/cases/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=geracao_de_leads</loc>
			<lastmod>2022-11-08T18:30:38+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/recurso-para-multas/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=exemplos_de_briefing</loc>
			<lastmod>2022-11-08T18:31:50+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/agencias-parceiras/programa/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=exemplos_de_briefing</loc>
			<lastmod>2022-11-08T18:31:51+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/recurso-para-multas/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=melhores_agencias_marketing_digital</loc>
			<lastmod>2022-11-08T18:32:06+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/recurso-para-multas/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=plano_negocios_agencia_marketing_digital</loc>
			<lastmod>2022-11-08T18:32:19+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/recurso-para-multas/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=nomes_para_agencia_de_marketing</loc>
			<lastmod>2022-11-08T18:32:20+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=equipe_de_marketing</loc>
			<lastmod>2022-11-08T18:32:21+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/</loc>
			<lastmod>2022-11-08T18:32:23+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=como_montar_agencia_marketing_digital</loc>
			<lastmod>2022-11-08T18:32:24+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/recurso-para-multas/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=como_montar_agencia_marketing_digital</loc>
			<lastmod>2022-11-08T18:32:24+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/recurso-para-multas/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=programas_de_parceria</loc>
			<lastmod>2022-11-08T18:32:24+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/agencias-parceiras/programa/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=programas_de_parceria</loc>
			<lastmod>2022-11-08T18:32:26+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/recurso-para-multas/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=ferramentas_agencia_de_marketing</loc>
			<lastmod>2022-11-08T18:32:29+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/recurso-para-multas/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=organograma_agencia_de_marketing</loc>
			<lastmod>2022-11-08T18:32:29+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=taboola</loc>
			<lastmod>2022-11-08T18:32:29+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=gestor_de_trafego</loc>
			<lastmod>2022-11-08T18:32:31+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=gerador_de_leads</loc>
			<lastmod>2022-11-08T18:32:31+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/botao-do-whatsapp/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=gerador_de_leads</loc>
			<lastmod>2022-11-08T18:32:31+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/formularios-para-geracao-de-leads/</loc>
			<lastmod>2022-11-08T18:32:32+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=puc_pr_post_link_whatsapp</loc>
			<lastmod>2022-11-08T18:32:35+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/gerador-de-link-whatsapp/</loc>
			<lastmod>2022-11-08T18:32:35+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/recurso-para-multas/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=como_captar_clientes_para_agencia_de_marketing</loc>
			<lastmod>2022-11-08T18:32:36+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-orcamentos-b2b/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=venda_consultiva</loc>
			<lastmod>2022-11-08T18:32:38+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=post_utm</loc>
			<lastmod>2022-11-08T18:32:39+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/botao-do-whatsapp/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=como_colocar_bot%C3%A3o_do_whatsapp_no_wix</loc>
			<lastmod>2022-11-08T18:32:40+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/recurso-para-multas/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=gestao_de_agencia_de_publicidade</loc>
			<lastmod>2022-11-08T18:32:40+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=tiktok_ads</loc>
			<lastmod>2022-11-08T18:32:42+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-leads/cases/cosiaco/</loc>
			<lastmod>2022-11-08T18:32:43+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=modelos_de_copywriting_prontas</loc>
			<lastmod>2022-11-08T18:32:44+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=blog_para_gerar_leads</loc>
			<lastmod>2022-11-08T18:32:45+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=blog_para_gerar_leads</loc>
			<lastmod>2022-11-08T18:32:46+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=copywriting_para_instagram</loc>
			<lastmod>2022-11-08T18:32:46+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=digitalks</loc>
			<lastmod>2022-11-08T18:32:47+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-orcamentos-b2b/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=retencao_de_clientes</loc>
			<lastmod>2022-11-08T18:32:50+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=copywriting_e_marketing_digital</loc>
			<lastmod>2022-11-08T18:32:51+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=chatbots_online</loc>
			<lastmod>2022-11-08T18:32:53+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=chatbot_para_wix</loc>
			<lastmod>2022-11-08T18:32:53+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=chatbot_para_wix</loc>
			<lastmod>2022-11-08T18:32:53+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=social_ads</loc>
			<lastmod>2022-11-08T18:33:01+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=melhores_canais_de_midia_paga</loc>
			<lastmod>2022-11-08T18:33:02+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=trafego_pago_e_organico</loc>
			<lastmod>2022-11-08T18:33:03+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=trafego_pago_organico</loc>
			<lastmod>2022-11-08T18:33:04+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=definir_objetivo_de_campanha</loc>
			<lastmod>2022-11-08T18:33:05+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=biblioteca_de_anuncios</loc>
			<lastmod>2022-11-08T18:33:05+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=biblioteca_de_anuncios</loc>
			<lastmod>2022-11-08T18:33:07+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=publico_alvo_anunciar_para_pessoas_certas</loc>
			<lastmod>2022-11-08T18:33:08+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=exemplos_de_anuncios</loc>
			<lastmod>2022-11-08T18:33:09+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=exemplos_de_titulos_criativos</loc>
			<lastmod>2022-11-08T18:33:10+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=plano_de_midia_paga</loc>
			<lastmod>2022-11-08T18:33:14+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=metricas_de_midia_paga</loc>
			<lastmod>2022-11-08T18:33:17+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=midia_paga_e_cro</loc>
			<lastmod>2022-11-08T18:33:18+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=midia_paga_e_cro</loc>
			<lastmod>2022-11-08T18:33:19+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=taxa_de_conversao_por_segmento</loc>
			<lastmod>2022-11-08T18:33:20+00:00</lastmod>
			<priority>0.51</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/materiais/checklist-para-geracao-e-nutricao-de-leads/</loc>
			<lastmod>2022-11-08T18:33:23+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=geracao_de_leads</loc>
			<lastmod>2022-11-08T18:33:24+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/botao-do-whatsapp/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=botao-whats</loc>
			<lastmod>2022-11-08T18:33:24+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=marketing_de_relacionamento</loc>
			<lastmod>2022-11-08T18:33:32+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=marketing_de_relacionamento</loc>
			<lastmod>2022-11-08T18:33:33+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=anuncios_online</loc>
			<lastmod>2022-11-08T18:33:34+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=anuncios_online</loc>
			<lastmod>2022-11-08T18:33:34+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=automacao_de_marketing</loc>
			<lastmod>2022-11-08T18:33:42+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-orcamentos-b2b/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=chatbot-para-vendas</loc>
			<lastmod>2022-11-08T18:33:42+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=chatbot-para-vendas</loc>
			<lastmod>2022-11-08T18:33:42+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=leads_para_energia_solar</loc>
			<lastmod>2022-11-08T18:33:43+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-orcamentos-b2b/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=como_gerar_leads_b2b</loc>
			<lastmod>2022-11-08T18:33:44+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=leads_para_corretores_de_seguros</loc>
			<lastmod>2022-11-08T18:33:45+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=leads_para_corretores_de_seguros</loc>
			<lastmod>2022-11-08T18:33:45+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=marketing_imobiliario</loc>
			<lastmod>2022-11-08T18:33:46+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=marketing_imobiliario</loc>
			<lastmod>2022-11-08T18:33:46+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=marketing_para_franquias</loc>
			<lastmod>2022-11-08T18:33:48+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=marketing_para_pequenas_empresas</loc>
			<lastmod>2022-11-08T18:33:49+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=marketing_juridico</loc>
			<lastmod>2022-11-08T18:33:51+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=marketing_para_construtoras</loc>
			<lastmod>2022-11-08T18:33:51+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=marketing_para_construtoras</loc>
			<lastmod>2022-11-08T18:33:51+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=google_analytics</loc>
			<lastmod>2022-11-08T18:33:52+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=chat_online</loc>
			<lastmod>2022-11-08T18:33:53+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=conversao_de_vendas</loc>
			<lastmod>2022-11-08T18:33:53+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-orcamentos-b2b/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=tomadores_de_decisao</loc>
			<lastmod>2022-11-08T18:33:55+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=como_anunciar_no_google_ads</loc>
			<lastmod>2022-11-08T18:33:59+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=como_anunciar_no_facebook_ads</loc>
			<lastmod>2022-11-08T18:33:59+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=funil_de_conversao</loc>
			<lastmod>2022-11-08T18:34:00+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/materiais/calculadora-de-geracao-de-leads/</loc>
			<lastmod>2022-11-08T18:34:02+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=melhorar_taxa_de_conversao</loc>
			<lastmod>2022-11-08T18:34:03+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=taxa_de_conversao</loc>
			<lastmod>2022-11-08T18:34:03+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=taxa_de_conversao</loc>
			<lastmod>2022-11-08T18:34:03+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=prova_social</loc>
			<lastmod>2022-11-08T18:34:03+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=landing_page_de_alta_conversao</loc>
			<lastmod>2022-11-08T18:34:04+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=caseblog</loc>
			<lastmod>2022-11-08T18:34:04+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/panorama-de-geracao-de-leads-no-brasil/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=como_se_preparar_para_a_black_friday</loc>
			<lastmod>2022-11-08T18:34:07+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=marketing_educacional</loc>
			<lastmod>2022-11-08T18:34:07+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=marketing_para_clinicas_de_estetica</loc>
			<lastmod>2022-11-08T18:34:08+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-orcamentos-b2b/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=marketing_industrial</loc>
			<lastmod>2022-11-08T18:34:08+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=custo_por_lead</loc>
			<lastmod>2022-11-08T18:34:08+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=captura_de_leads</loc>
			<lastmod>2022-11-08T18:34:12+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/botao-do-whatsapp/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=como_colocar_botao_gratis</loc>
			<lastmod>2022-11-08T18:34:13+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=melhores_chatbots</loc>
			<lastmod>2022-11-08T18:34:13+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=como_criar_chatbot_gratis</loc>
			<lastmod>2022-11-08T18:34:14+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=tipos_de_chatbots</loc>
			<lastmod>2022-11-08T18:34:14+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-orcamentos-b2b/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=tipos_de_chatbots</loc>
			<lastmod>2022-11-08T18:34:15+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/botao-do-whatsapp/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=texto_pronto_para_vendas_no_whatsapp</loc>
			<lastmod>2022-11-08T18:34:17+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=rede_de_pesquisa_google</loc>
			<lastmod>2022-11-08T18:34:19+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=google_ads_vs_facebook_ads</loc>
			<lastmod>2022-11-08T18:34:20+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=trafego_para_site</loc>
			<lastmod>2022-11-08T18:34:20+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=pixel_do_facebook</loc>
			<lastmod>2022-11-08T18:34:20+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=motivos_para_investir_em_anuncios</loc>
			<lastmod>2022-11-08T18:34:21+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=anuncios_de_black_friday</loc>
			<lastmod>2022-11-08T18:34:21+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=trafego_pago</loc>
			<lastmod>2022-11-08T18:34:22+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=trafego_pago_geracao_de_leads</loc>
			<lastmod>2022-11-08T18:34:22+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=estrategia_de_marketing_conversacional</loc>
			<lastmod>2022-11-08T18:34:22+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/recurso-para-multas/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=plano_de_acao</loc>
			<lastmod>2022-11-08T18:34:25+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=campanha_publicitaria</loc>
			<lastmod>2022-11-08T18:34:26+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/estudo-de-caso/recurso-para-multas/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=campanha_publicitaria</loc>
			<lastmod>2022-11-08T18:34:26+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/botao-do-whatsapp/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=vendas_pelo_whatsapp</loc>
			<lastmod>2022-11-08T18:34:26+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=formularios_para_captacao_de_leads</loc>
			<lastmod>2022-11-08T18:34:34+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/contato/?utm_content=ai-contact-center&amp;utm_term=ai-cc&amp;utm_medium=button&amp;utm_source=blog&amp;utm_campaign=Inteligencia%20em%20contact%20center</loc>
			<lastmod>2022-11-08T18:34:35+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=storytelling</loc>
			<lastmod>2022-11-08T18:34:38+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/geracao-de-orcamentos-b2b/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=inbound_ou_outbound</loc>
			<lastmod>2022-11-08T18:34:38+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=puc_pr</loc>
			<lastmod>2022-11-08T18:34:38+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/chatbot-leadster/?utm_source=materiais&amp;utm_campaign=whatsapp_gerador</loc>
			<lastmod>2022-11-08T18:34:39+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=crm_de_vendas</loc>
			<lastmod>2022-11-08T18:34:39+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=planejamento_estrategico_de_marketing</loc>
			<lastmod>2022-11-08T18:34:40+00:00</lastmod>
			<priority>0.41</priority>
		</url>
		<url>
			<loc>https://leadster.com.br/leads-qualificados-com-anuncios/?utm_source=blog&amp;utm_medium=post&amp;utm_campaign=copy_para_vendas</loc>
			<lastmod>2022-11-08T18:34:40+00:00</lastmod>
			<priority>0.41</priority>
		</url>
	</urlset>
	`))

	siteMap := site_map.NewSiteMap("http://www.teste.com")
	err := siteMap.Call()
	log.Println(err)
	assert.Nil(t, err)
}
