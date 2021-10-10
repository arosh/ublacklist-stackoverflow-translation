require 'yaml'

puts "[uBlock Origin]"
puts "! Title: uBlacklist Stack Overflow Translation"
puts "! Description: Filters that enable excluding the machine-translated sites of Stack Overflow in Firefox for Android"
puts "! Expires: 1 day (update frequency)"
puts "! Homepage: https://github.com/arosh/ublacklist-stackoverflow-translation"
puts "! License: CC0 1.0 (https://creativecommons.org/publicdomain/zero/1.0/)"
puts
puts YAML.load(STDIN).map {
  host = _1["domain"].gsub(/^\*\./, "")
  if _1["domain"].start_with?("*.") or _1["evidence"].start_with?("http://") then
    "www.google.*##.kCrYT > a[href*=\"#{host}/\"]:upward(.xpd)\n" + # Firefox default
    "www.google.*##.C8nzq[href*=\"#{host}/\"]:upward(.xpd)\n" + # Google Search Fixer enabled
    "www.google.*##.xpdopen .sXtWJb[href*=\"#{host}/\"]:upward(.xpdopen)\n" + # Featured snippets (Google Search Fixer)
    "www.google.*##.aI1xUe .sXtWJb[href*=\"#{host}/\"]:upward(.aI1xUe)\n" + # People also ask (Google Search Fixer)
    "cn.bing.com,www.bing.com##.b_algoheader > a[href*=\"#{host}/\"]:upward(.b_algo)\n" +
    "duckduckgo.com##div[data-domain$=\"#{host}\"]\n" +
    "www.ecosia.org##.result[data-url*=\"#{host}/\"]\n" +
    "startpage.com##.w-gl__result-title[href*=\"#{host}/\"]:upward(F.w-gl__result)\n" +
    # Images
    "www.google.*##.isv-r[data-ru*=\"#{host}/\"]\n" + # Firefox default
    "www.google.*##.VFACy[href*=\"#{host}/\"]:upward(.isv-r)\n" + # Google Search Fixer enabled
    "cn.bing.com,www.bing.com##a[title$=\"#{host}\"]:upward(.dgControl_list > li)\n" +
    "duckduckgo.com##.tile--img__sub[href*=\"#{host}/\"]:upward(.tile--img)\n"
  else
    "www.google.*##.kCrYT > a[href^=\"/url?q=https://#{host}/\"]:upward(.xpd)\n" + # Firefox default
    "www.google.*##.C8nzq[href^=\"https://#{host}/\"]:upward(.xpd)\n" + # Google Search Fixer enabled
    "www.google.*##.xpdopen .sXtWJb[href^=\"https://#{host}/\"]:upward(.xpdopen)\n" + # Featured snippets (Google Search Fixer)
    "www.google.*##.aI1xUe .sXtWJb[href^=\"https://#{host}/\"]:upward(.aI1xUe)\n" + # People also ask (Google Search Fixer)
    "cn.bing.com,www.bing.com##.b_algoheader > a[href^=\"https://#{host}/\"]:upward(.b_algo)\n" +
    "duckduckgo.com##div[data-domain$=\"#{host}\"]\n" +
    "www.ecosia.org##.result[data-url^=\"https://#{host}/\"]\n" +
    "startpage.com##.w-gl__result-title[href^=\"https://#{host}/\"]:upward(.w-gl__result)\n" +
    # Images
    "www.google.*##.isv-r[data-ru^=\"https://#{host}/\"]\n" + # Firefox default
    "www.google.*##.VFACy[href^=\"https://#{host}/\"]:upward(.isv-r)\n" + # Google Search Fixer enabled
    "cn.bing.com,www.bing.com##.inflnk[href^=\"https://#{host}/\"]:upward(.dgControl_list > li)\n" +
    "duckduckgo.com##.tile--img__sub[href^=\"https://#{host}/\"]:upward(.tile--img)\n"
  end
}
