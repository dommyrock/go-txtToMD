package static

var HtmlOpen = `<!doctype html>
<html>

<head>
   <meta charset="utf-8">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <title>HTML document</title>
   <style>
      :root {
         padding-left: 17rem;
         padding-right: 17rem;
         padding-top: 0.5rem;
         height: 100vh;
			background: #f1f1f1;
      }
		
		body {
			background: rgb(198 198 198 / 17%);
			border-radius: 10px;
    		padding: 0.1rem;
			padding-left: 2rem;
			padding-right: 2rem;
			font-family: -apple-system,BlinkMacSystemFont,"Segoe UI",Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji";
		}


      blockquote {
         font-style: italic;
         text-align: left;
         background-color: rgb(239 236 236 / 72%);
			border-left: 0.2rem solid #70c3d6;
			padding-left: 0.5rem;
      }

      table {
         font-family: Arial, Helvetica, sans-serif;
         border-collapse: collapse;
         text-align: center;
      }

      table th:first-child {
         border-radius: 5px 0 0 0;
      }

      table th:last-child {
         border-radius: 0 5px 0 0;
      }

      table td,
      table th {
         border: 1px solid #ff000003;
         padding: 8px;
      }

      table tr:nth-child(even) {
         background-color: #f2f2f2;
      }

      table tr:hover {
         background-color: #ddd;
      }

      table th {
         padding-top: 12px;
         padding-bottom: 12px;
         text-align: center;
         background-color: #70c3d6;
         color: white;
      }

      table tbody tr:last-of-type {
         border-bottom: 2px solid #81e0cc;
      }

      ul {
         padding: 0.1rem;
      }

      ul li {
         /* remove bulet icon */
         display: inline-block;
         background-color: rgb(204, 222, 240);
         border: solid 4px #ccc;
      }
      li{
         transition: transform .3s ease-out;
      }
      ul>li:hover{
         transform: translate(20px, 0);
      }

      a {
         text-decoration: none;
         box-shadow: inset 0 -2px 0 rgba(123, 228, 163, 0.5), 0 2px 0 rgba(123, 228, 163, 0.5);
         transition: box-shadow .3s;
         color: inherit;
         overflow: hidden;
      }

      a:hover {
         box-shadow: inset 0 -30px 0 rgba(123, 228, 163, 0.5), 0 2px 0 rgba(123, 228, 163, 0.5);
      }

		/* Background */ .bg { color: #f8f8f2; background-color: #282a36 }
		/* PreWrapper */ .chroma { color: #f8f8f2; background-color: #282a36; border-radius: 4px; padding: 0.5rem; }
		/* LineTableTD */ .chroma .lntd { vertical-align: top; padding: 0; margin: 0; border: 0; }
		/* LineTable */ .chroma .lntable { border-spacing: 0; padding: 0; margin: 0; border: 0; }
		/* LineHighlight */ .chroma .hl { background-color: #3d3f4a }
		/* LineNumbersTable */ .chroma .lnt { white-space: pre; user-select: none; margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
		/* LineNumbers */ .chroma .ln { white-space: pre; user-select: none; margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
		/* Line */ .chroma .line { display: flex; }
		/* Keyword */ .chroma .k { color: #ff79c6 }
		/* KeywordConstant */ .chroma .kc { color: #ff79c6 }
		/* KeywordDeclaration */ .chroma .kd { color: #8be9fd; font-style: italic }
		/* KeywordNamespace */ .chroma .kn { color: #ff79c6 }
		/* KeywordPseudo */ .chroma .kp { color: #ff79c6 }
		/* KeywordReserved */ .chroma .kr { color: #ff79c6 }
		/* KeywordType */ .chroma .kt { color: #8be9fd }
		/* NameAttribute */ .chroma .na { color: #50fa7b }
		/* NameBuiltin */ .chroma .nb { color: #8be9fd; font-style: italic }
		/* NameClass */ .chroma .nc { color: #50fa7b }
		/* NameFunction */ .chroma .nf { color: #50fa7b }
		/* NameLabel */ .chroma .nl { color: #8be9fd; font-style: italic }
		/* NameTag */ .chroma .nt { color: #ff79c6 }
		/* NameVariable */ .chroma .nv { color: #8be9fd; font-style: italic }
		/* NameVariableClass */ .chroma .vc { color: #8be9fd; font-style: italic }
		/* NameVariableGlobal */ .chroma .vg { color: #8be9fd; font-style: italic }
		/* NameVariableInstance */ .chroma .vi { color: #8be9fd; font-style: italic }
		/* LiteralString */ .chroma .s { color: #f1fa8c }
		/* LiteralStringAffix */ .chroma .sa { color: #f1fa8c }
		/* LiteralStringBacktick */ .chroma .sb { color: #f1fa8c }
		/* LiteralStringChar */ .chroma .sc { color: #f1fa8c }
		/* LiteralStringDelimiter */ .chroma .dl { color: #f1fa8c }
		/* LiteralStringDoc */ .chroma .sd { color: #f1fa8c }
		/* LiteralStringDouble */ .chroma .s2 { color: #f1fa8c }
		/* LiteralStringEscape */ .chroma .se { color: #f1fa8c }
		/* LiteralStringHeredoc */ .chroma .sh { color: #f1fa8c }
		/* LiteralStringInterpol */ .chroma .si { color: #f1fa8c }
		/* LiteralStringOther */ .chroma .sx { color: #f1fa8c }
		/* LiteralStringRegex */ .chroma .sr { color: #f1fa8c }
		/* LiteralStringSingle */ .chroma .s1 { color: #f1fa8c }
		/* LiteralStringSymbol */ .chroma .ss { color: #f1fa8c }
		/* LiteralNumber */ .chroma .m { color: #bd93f9 }
		/* LiteralNumberBin */ .chroma .mb { color: #bd93f9 }
		/* LiteralNumberFloat */ .chroma .mf { color: #bd93f9 }
		/* LiteralNumberHex */ .chroma .mh { color: #bd93f9 }
		/* LiteralNumberInteger */ .chroma .mi { color: #bd93f9 }
		/* LiteralNumberIntegerLong */ .chroma .il { color: #bd93f9 }
		/* LiteralNumberOct */ .chroma .mo { color: #bd93f9 }
		/* Operator */ .chroma .o { color: #ff79c6 }
		/* OperatorWord */ .chroma .ow { color: #ff79c6 }
		/* Comment */ .chroma .c { color: #6272a4 }
		/* CommentHashbang */ .chroma .ch { color: #6272a4 }
		/* CommentMultiline */ .chroma .cm { color: #6272a4 }
		/* CommentSingle */ .chroma .c1 { color: #6272a4 }
		/* CommentSpecial */ .chroma .cs { color: #6272a4 }
		/* CommentPreproc */ .chroma .cp { color: #ff79c6 }
		/* CommentPreprocFile */ .chroma .cpf { color: #ff79c6 }
		/* GenericDeleted */ .chroma .gd { color: #ff5555 }
		/* GenericEmph */ .chroma .ge { text-decoration: underline }
		/* GenericHeading */ .chroma .gh { font-weight: bold }
		/* GenericInserted */ .chroma .gi { color: #50fa7b; font-weight: bold }
		/* GenericOutput */ .chroma .go { color: #44475a }
		/* GenericSubheading */ .chroma .gu { font-weight: bold }
		/* GenericUnderline */ .chroma .gl { text-decoration: underline }
   </style>
</head>

<body>
<div id="content">`
var HtmlEnd = "</div>\n</body>\n</html>"
