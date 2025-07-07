#!/bin/sh

echo "# Отчет о нагрузочном тестировании. Lab_05." > load_test_report.md

echo "#### ab -n 25000 -c 10 http://localhost/api/v1/competitions" >> load_test_report.md

echo "\`\`\`" >>  load_test_report.md

ab -n 25000 -c 10 http://localhost/api/v1/competitions >> load_test_report.md

echo "\`\`\`" >>  load_test_report.md

echo "#### ab -n 25000 -c 10 http://localhost/api/v1/no_balance/competitions" >> load_test_report.md

echo "\`\`\`" >>  load_test_report.md

ab -n 25000 -c 10 http://localhost/api/v1/no_balance/competitions >> load_test_report.md

echo "\`\`\`" >>  load_test_report.md