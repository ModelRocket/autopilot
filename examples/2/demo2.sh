#!/bin/sh

GREEN='\033[0;32m'
BOLD_GREEN='\033[3;33m'
RED='\033[3;31m'
YELLOW='\033[0;33m'
NC='\033[0m'

clear
cat banner2
sleep 5
clear

printf "root@kubemaster# cat autopilot.yaml\n"
cat autopilot.yaml
sleep 5
clear
printf "${GREEN}root@kubemaster# kubectl apply -f autopilot.sh${NC}\n"
sleep 1
echo "configmap "autopilot-config" created"
sleep 1
echo "serviceaccount \"autopilot-account\" created"
sleep 1
echo "clusterrole \"autopilot-role\" created"
sleep 1
echo "clusterrolebinding \"autopilot-role-binding\" created"
sleep 1
echo "deployment \"autopilot\" created"
sleep 5


clear
printf "root@kubemaster# cat policy-example.yaml\n"
cat policy-example-2.yaml
sleep 1
printf "${GREEN}root@kubemaster# kubectl apply -f policy-example.yaml${NC}\n"
sleep 1
echo "storagepolicy "volume-fix-latency-policy" created"
sleep 5

# DUMP PVC BEFORE
clear
printf "root@kubemaster# kubectl describe pvc px-postgres-pvc | grep Nodes\n"
# cat pvc-before
printf "${BOLD_GREEN}Nodes:      worker2${NC}"
sleep 5
clear

# Get AutoPilot output
printf "root@kubemaster# kubectl get events --field-selector involvedObject.name=volume-fix-latency-policy  -o=custom-columns=Type:.type,Reason:.reason,Message:.message\n"
sleep 1
echo "TYPE     REASON                  MESSAGE"
echo "Normal   MonitoringStarted       Started monitoring volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7"
sleep 2

echo
printf "root@kubemaster# kubectl get events --field-selector involvedObject.name=volume-fix-latency-policy  -o=custom-columns=Type:.type,Reason:.reason,Message:.message\n"
sleep 1
echo "TYPE     REASON                  MESSAGE"
echo "Normal   MonitoringStarted       Started monitoring volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7"
printf "${YELLOW}Warning  ConditonApproaching     Volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7 has met condition: volume_latency_ms > 15${NC}\n"
sleep 2

echo 
printf "root@kubemaster# kubectl get events --field-selector involvedObject.name=volume-fix-latency-policy  -o=custom-columns=Type:.type,Reason:.reason,Message:.message\n"
sleep 1
echo "TYPE     REASON                  MESSAGE"
echo "Normal   MonitoringStarted       Started monitoring volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7"
printf "${YELLOW}Warning  ConditonApproaching     Volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7 has met condition: volume_latency_ms > 15${NC}\n"
printf "${YELLOW}Warning  ConditonApproaching     Volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7 has met condition: volume_latency_ms > 20${NC}\n"
sleep 2

echo 
printf "root@kubemaster# kubectl get events --field-selector involvedObject.name=volume-fix-latency-policy  -o=custom-columns=Type:.type,Reason:.reason,Message:.message\n"
sleep 1
echo "TYPE     REASON                  MESSAGE"
echo "Normal   MonitoringStarted       Started monitoring volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7"
printf "${YELLOW}Warning  ConditonApproaching     Volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7 has met condition: volume_latency_ms > 15${NC}\n"
printf "${YELLOW}Warning  ConditonApproaching     Volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7 has met condition: volume_latency_ms > 20${NC}\n"
printf "${RED}Warning  ConditonMet             Volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7 has met condition: volume_latency_ms > 25${NC}\n"
sleep 2

echo 
printf "root@kubemaster# kubectl get events --field-selector involvedObject.name=volume-fix-latency-policy  -o=custom-columns=Type:.type,Reason:.reason,Message:.message\n"
sleep 1
echo "TYPE     REASON                  MESSAGE"
echo "Normal   MonitoringStarted       Started monitoring volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7"
printf "${YELLOW}Warning  ConditonApproaching     Volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7 has met condition: volume_latency_ms > 15${NC}\n"
printf "${YELLOW}Warning  ConditonApproaching     Volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7 has met condition: volume_latency_ms > 20${NC}\n"
printf "${RED}Warning  ConditonMet             Volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7 has met condition: volume_latency_ms > 25${NC}\n"
printf "${RED}Warning  ActionTriggered         Action: evict triggered for volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7. Evict from: worker2${NC}\n"
sleep 2

echo 
printf "root@kubemaster# kubectl get events --field-selector involvedObject.name=volume-fix-latency-policy  -o=custom-columns=Type:.type,Reason:.reason,Message:.message\n"
sleep 1
echo "TYPE     REASON                  MESSAGE"
echo "Normal   MonitoringStarted       Started monitoring volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7"
printf "${YELLOW}Warning  ConditonApproaching     Volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7 has met condition: volume_latency_ms > 15${NC}\n"
printf "${YELLOW}Warning  ConditonApproaching     Volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7 has met condition: volume_latency_ms > 20${NC}\n"
printf "${RED}Warning  ConditonMet             Volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7 has met condition: volume_latency_ms > 25${NC}\n"
printf "${RED}Warning  ActionTriggered         Action: evict triggered for volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7. Evict from: worker2${NC}\n"
printf "${GREEN}Normal   ActionSuccessful        Volume pvc-82ced888-19f6-11e9-a9a4-080027ee1df7 has been evicted from: worker2 and moved to: worker3${NC}\n"
sleep 6


# DUMP PVC AFTER
clear
printf "root@kubemaster# kubectl describe pvc px-postgres-pvc | grep Nodes\n"
# cat pvc-after
printf "${BOLD_GREEN}Nodes:      worker3${NC}"
sleep 5
echo
sleep 5
