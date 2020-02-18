DATE=`date`
RED='\033[0;31m'
GREEN='\033[0;92m'
ORANGE='\033[38;2;255;165;0m'
BLUE='\033[0;94m'
NC='\033[0m'
COMMIT_MSG=""

if [ -z $1 ]; then
    COMMIT_MSG=$1
else
    COMMIT_MSG="$DATE - auto commit"
fi

printMessage() {
    printf '=%.0s' {1..80}; echo
    echo -e "$1 $2"
    printf '=%.0s' {1..80}; echo
}

handleError () {
    if [ $1 -eq 1 ]; then
        printMessage "${RED}ERROR:${NC} $2 $3"
        exit 1
    else
        printMessage "${GREEN}SUCCESS:${NC} All checks passed for $3"
    fi
}

printMessage "${BLUE}INFO:${NC} Starting post operation"

./godelw format
printMessage "${BLUE}INFO:${NC} Formatted code files"

./godelw license
printMessage "${BLUE}INFO:${NC} Added license headers"

./godelw check
handleError $? "Can't post due to errors. Please review" "${ORANGE}./godelw check${NC}"

./godelw test
handleError $? "Can't post due to errors. Please review" "${ORANGE}./godelw test${NC}"

printMessage "${GREEN}SUCCESS:${NC} All checks passed"

printMessage "${BLUE}INFO:${NC} Generating auto-commit: ${ORANGE}${COMMIT_MSG}${NC}"
git add .
git config --global user.name "Kelcey Jamison-Damage"
git commit -m "${COMMIT_MSG}"

printMessage "${BLUE}INFO:${NC} Pushing to Git"
git push origin master
handleError $? "Can't post due to errors. Please review" "${ORANGE}git push origin master${NC}"

