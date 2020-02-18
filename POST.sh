RED='\033[0;31m'
GREEN='\033[0;92m'
ORANGE='\033[38;2;255;165;0m'
NC='\033[0m'

printMessage() {
    printf '=%.0s' {1..80}; echo
    echo -e $1
    printf '=%.0s' {1..80}; echo
}

handleError () {
    if [ $1 -eq 1 ]; then
        printMessage "${RED}ERROR:${NC} $2"
        exit 1
    fi
}

DATE=`date`
./godelw format
./godelw license
./godelw check
handleError $? "Can't post due to errors. Please review ${ORANGE}./godelw check${NC}"
./godelw test
handleError $? "Can't post due to errors. Please review ${ORANGE}./godelw test${NC}"
printMessage "${GRREN}SUCCESS:${NC} All checks passed"
git add .
git config --global user.name "Kelcey Jamison-Damage"
git commit -m "$DATE - auto commit"
git push origin master

