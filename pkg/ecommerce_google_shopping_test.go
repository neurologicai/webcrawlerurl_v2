package pkg_test

import (
	"testing"

	"github.com/gildemberg-santos/webcrawlerurl_v2/pkg"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestEcommerceGoogleShopping_Call(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://www.teste.com/google_shopping.xml", httpmock.NewStringResponder(200, `
	<?xml version="1.0" encoding="UTF-8"?>
	<feed xmlns="http://www.w3.org/2005/Atom">
		<entry>
			<g:link>https://www.google.com/shopping/product/1</g:link>
		</entry>
		<entry>
			<g:link>https://www.google.com/shopping/product/2</g:link>
		</entry>
	</feed>
	`))

	ecommerceGoogleShopping := pkg.NewEcommerceGoogleShopping("http://www.teste.com/google_shopping.xml", "https://www.google.com/shopping/product/**", 30)
	response := ecommerceGoogleShopping.Call()

	assert.NotNil(t, response)
	assert.Equal(t, "http://www.teste.com/google_shopping.xml", ecommerceGoogleShopping.Url)
	assert.Len(t, ecommerceGoogleShopping.Urls, 2)
	assert.Equal(t, "https://www.google.com/shopping/product/1", ecommerceGoogleShopping.Urls[0])
	assert.Equal(t, "https://www.google.com/shopping/product/2", ecommerceGoogleShopping.Urls[1])
}
