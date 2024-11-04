Severity	Age (Days)
MEDIUM	31
MEDIUM	38
HIGH	38
HIGH	5
HIGH	38
MEDIUM	31
HIGH	5

当前路径的data.txt里有如上一组数据，左边是漏洞定级，右边是漏洞age，CRITICAL级别的漏洞要3个自然日内修复，HIGH级别的漏洞要7个自然日内修复，MEDIUM要10个自然日内修复，LOW要20个自然日内修复。此程序读取data.txt，计算按期修复率，并打印每一组数据是否有按期修复，如：HIGH	5[按期修复]

