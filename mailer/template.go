package mailer

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/dmitrymomot/sprig"
)

type mailTemplate struct {
	UserID  string
	Name    string
	Subject string

	Preheader string
	Intro     []string
	Outro     []string
	Button    *buttonTpl
	Footer    []string

	UnsubscribeLink string
	RemoveEmailLink string

	Product struct {
		Name string
		Link string
	}

	Email string
	code  string
}

// ButtonTpl button template element
type buttonTpl struct {
	Link  string
	Title string
}

func parseTemplate(mail Mailer, mailTpl *mailTemplate) (string, error) {
	baseTpl, err := mail.box().MustString(mail.config().BaseTemplate())
	if err != nil {
		return "", err
	}

	if mailTpl.Product.Name == "" {
		mailTpl.Product.Name = mail.config().ProductName()
	}
	if mailTpl.Product.Link == "" {
		mailTpl.Product.Link = mail.config().ProductLink()
	}

	t := template.Must(template.New("").Funcs(sprig.FuncMap()).Parse(baseTpl))
	var buffer bytes.Buffer
	if err = t.Execute(&buffer, mailTpl); err != nil {
		return "", err
	}

	buffStr := replacePlaceholders(buffer.String(), mailTpl)

	return buffStr, nil
}

func replacePlaceholders(s string, mailTpl *mailTemplate) string {
	s = strings.Replace(s, ":name", mailTpl.Name, -1)
	s = strings.Replace(s, ":user_id", mailTpl.UserID, -1)
	s = strings.Replace(s, ":email", mailTpl.Email, -1)
	s = strings.Replace(s, ":confirmation_code", mailTpl.code, -1)
	s = strings.Replace(s, ":product_name", mailTpl.Product.Name, -1)
	s = strings.Replace(s, ":product_link", mailTpl.Product.Link, -1)
	return s
}

func getDefaultTemplateHTML() string {
	return `
	<!DOCTYPE html>
	<html>
	  <head>
	    <meta name="viewport" content="width=device-width">
	    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
	    <title>{{ .Subject }}</title>
	    <style>
	    /* -------------------------------------
	        INLINED WITH htmlemail.io/inline
	    ------------------------------------- */
	    /* -------------------------------------
	        RESPONSIVE AND MOBILE FRIENDLY STYLES
	    ------------------------------------- */
	    @media only screen and (max-width: 620px) {
	      table[class=body] h1 {
	        font-size: 28px !important;
	        margin-bottom: 10px !important;
	      }
	      table[class=body] p,
	            table[class=body] ul,
	            table[class=body] ol,
	            table[class=body] td,
	            table[class=body] span,
	            table[class=body] a {
	        font-size: 16px !important;
	      }
	      table[class=body] .wrapper,
	            table[class=body] .article {
	        padding: 10px !important;
	      }
	      table[class=body] .content {
	        padding: 0 !important;
	      }
	      table[class=body] .container {
	        padding: 0 !important;
	        width: 100% !important;
	      }
	      table[class=body] .main {
	        border-left-width: 0 !important;
	        border-radius: 0 !important;
	        border-right-width: 0 !important;
	      }
	      table[class=body] .btn table {
	        width: 100% !important;
	      }
	      table[class=body] .btn a {
	        width: 100% !important;
	      }
	      table[class=body] .img-responsive {
	        height: auto !important;
	        max-width: 100% !important;
	        width: auto !important;
	      }
	    }

	    /* -------------------------------------
	        PRESERVE THESE STYLES IN THE HEAD
	    ------------------------------------- */
	    @media all {
	      .ExternalClass {
	        width: 100%;
	      }
	      .ExternalClass,
	            .ExternalClass p,
	            .ExternalClass span,
	            .ExternalClass font,
	            .ExternalClass td,
	            .ExternalClass div {
	        line-height: 100%;
	      }
	      .apple-link a {
	        color: inherit !important;
	        font-family: inherit !important;
	        font-size: inherit !important;
	        font-weight: inherit !important;
	        line-height: inherit !important;
	        text-decoration: none !important;
	      }
	      .btn-primary table td:hover {
	        background-color: #34495e !important;
	      }
	      .btn-primary a:hover {
	        background-color: #34495e !important;
	        border-color: #34495e !important;
	      }
	    }
	    </style>
	  </head>
	  <body class="" style="background-color: #f6f6f6; font-family: sans-serif; -webkit-font-smoothing: antialiased; font-size: 14px; line-height: 1.4; margin: 0; padding: 0; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%;">
	    <table border="0" cellpadding="0" cellspacing="0" class="body" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; background-color: #f6f6f6;">
	      <tr>
	        <td style="font-family: sans-serif; font-size: 14px; vertical-align: top;">&nbsp;</td>
	        <td class="container" style="font-family: sans-serif; font-size: 14px; vertical-align: top; display: block; Margin: 0 auto; max-width: 580px; padding: 10px; width: 580px;">
	          <div class="content" style="box-sizing: border-box; display: block; Margin: 0 auto; max-width: 580px; padding: 10px;">

	            <!-- START CENTERED WHITE CONTAINER -->
	            {{ if .Preheader }}
	            <span class="preheader" style="color: transparent; display: none; height: 0; max-height: 0; max-width: 0; opacity: 0; overflow: hidden; mso-hide: all; visibility: hidden; width: 0;">{{ .Preheader }}</span>
	            {{ end }}

	            <table class="main" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; background: #ffffff; border-radius: 3px;">

	              <!-- START MAIN CONTENT AREA -->
	              <tr>
	                <td class="wrapper" style="font-family: sans-serif; font-size: 14px; vertical-align: top; box-sizing: border-box; padding: 20px;">
	                  <table border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%;">
	                    <tr>
	                      <td style="font-family: sans-serif; font-size: 14px; vertical-align: top;">
	                        <p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; Margin-bottom: 15px;">Hi{{ if .Name }}&nbsp;{{ .Name }}{{ end }},</p>
	                        {{ range .Intro }}
	                        <p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; Margin-bottom: 15px;">{{ . }}</p>
	                        {{ end }}
	                        {{ if .Button }}
	                        <table border="0" cellpadding="0" cellspacing="0" class="btn btn-primary" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%; box-sizing: border-box;">
	                          <tbody>
	                            <tr>
	                              <td align="center" style="font-family: sans-serif; font-size: 14px; vertical-align: top; padding-bottom: 15px;">
	                                <table border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: auto;">
	                                  <tbody>
	                                    <tr>
	                                      <td style="font-family: sans-serif; font-size: 14px; vertical-align: top; background-color: #3498db; border-radius: 5px; text-align: center;"> <a href="{{ .Button.Link }}" target="_blank" style="display: inline-block; color: #ffffff; background-color: #3498db; border: solid 1px #3498db; border-radius: 5px; box-sizing: border-box; cursor: pointer; text-decoration: none; font-size: 14px; font-weight: bold; margin: 0; padding: 12px 25px; text-transform: capitalize; border-color: #3498db;">{{ .Button.Title }}</a> </td>
	                                    </tr>
	                                  </tbody>
	                                </table>
	                              </td>
	                            </tr>
	                          </tbody>
	                        </table>
	                        {{ end }}
	                        {{ range .Outro }}
	                        <p style="font-family: sans-serif; font-size: 14px; font-weight: normal; margin: 0; Margin-bottom: 15px;">{{ . }}</p>
	                        {{ end }}
	                      </td>
	                    </tr>
	                  </table>
	                </td>
	              </tr>

	            <!-- END MAIN CONTENT AREA -->
	            </table>

	            <!-- START FOOTER -->
	            <div class="footer" style="clear: both; Margin-top: 10px; text-align: center; width: 100%;">
	              <table border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%;">
	                <tr>
	                  <td class="content-block" style="font-family: sans-serif; vertical-align: top; padding-bottom: 10px; padding-top: 10px; font-size: 12px; color: #999999; text-align: center;">
	                    {{ range .Footer }}
	                    <span class="apple-link" style="color: #999999; font-size: 12px; text-align: center;">{{ . }}</span>
	                    {{ end }}
	                    {{ if .Links.Unsubscribe }}
	                    <br>Don't like these emails? <a target="_blank" href="{{ .Links.Unsubscribe }}" style="text-decoration: underline; color: #999999; font-size: 12px; text-align: center;">Unsubscribe</a>.
	                    {{ end }}
	                    {{ if .Links.RemoveEmail }}
	                    <br>If this is not your account you can <a target="_blank" href="{{ .Links.RemoveEmail }}" style="text-decoration: underline; color: #999999; font-size: 12px; text-align: center;">remove your email from it.</a>.
	                    {{ end }}
	                  </td>
	                </tr>
	                {{ if .Product }}
	                <tr>
	                  <td class="content-block powered-by" style="font-family: sans-serif; vertical-align: top; padding-bottom: 10px; padding-top: 10px; font-size: 12px; color: #999999; text-align: center;">
	                    &copy; {{ now.Year }} <a target="_blank" href="{{ .Product.Link }}" style="color: #999999; font-size: 12px; text-align: center; text-decoration: none;">{{ .Product.Name }}</a>. All rights reserved.
	                  </td>
	                </tr>
	                {{ end }}
	              </table>
	            </div>
	            <!-- END FOOTER -->

	          <!-- END CENTERED WHITE CONTAINER -->
	          </div>
	        </td>
	        <td style="font-family: sans-serif; font-size: 14px; vertical-align: top;">&nbsp;</td>
	      </tr>
	    </table>
	  </body>
	</html>
	`
}
