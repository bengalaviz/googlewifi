#!/bin/bash

echo "Run the Google Wifi diagnostic-report and pull the infoJSON tag"
googlewifidump | jq -r '.infoJSON' > info.json

# This is a hacky way to fix the JSON output so that we can parse the info.
echo "Reading info.json"
readarray -t LINES < info.json

TEMP_OUTPUT_FILE=temp.json
FINAL_OUTPUT_FILE=final.json

echo "{" > ${TEMP_OUTPUT_FILE}
for LINE in "${LINES[@]}"; do
	TRIM_SPACES="$(echo -e "${LINE}" | tr -d '[:space:]')"
	i=$((${#TRIM_SPACES}-1))
	LAST_CHAR=`echo "${TRIM_SPACES:$i:1}"`
	if [[ "${LAST_CHAR}" == "{" ]]; then
		# Output is missing ":" before the {
		FIXED_LINE="${TRIM_SPACES::-1} : {"
		# This is were the hack is. Need to put [] around "station_state_updates" for it 
		# to parse correctly.
		if [[ "${FIXED_LINE}" == "station_state_updates : {" ]]; then
			echo "station_state_updates : [{" >> ${TEMP_OUTPUT_FILE}
		elif [[ "${FIXED_LINE}" == "station_state_update : {" ]]; then
			echo "},{${FIXED_LINE}" >> ${TEMP_OUTPUT_FILE}
		elif [[ "${FIXED_LINE}" == "port_forwardings : {" ]]; then
			echo "]${FIXED_LINE}" >> ${TEMP_OUTPUT_FILE}
		else
			echo "${FIXED_LINE}" >> ${TEMP_OUTPUT_FILE}
		fi
	else
		echo "${LINE}" >> ${TEMP_OUTPUT_FILE}
	fi
done
echo "}" >> ${TEMP_OUTPUT_FILE}

# Another hack to remove "\" character
sed 's/\\/\ /g' ${TEMP_OUTPUT_FILE} > ${FINAL_OUTPUT_FILE}
rm -fr ${TEMP_OUTPUT_FILE}
