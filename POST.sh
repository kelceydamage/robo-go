# System defaults
#================================================================================
DATE=`date`
RED='\033[0;31m'
GREEN='\033[0;92m'
ORANGE='\033[38;2;255;165;0m'
BLUE='\033[0;94m'
NC='\033[0m'
COMMIT_MSG=""
#================================================================================

# Git configuration
#================================================================================
GIT_USER="Kelcey Jamison-Damage"
GIT_BRANCH=`git rev-parse --abbrev-ref HEAD`
GIT_UPSTREAM=`git remote`
GIT_AUTOCOMMIT_MSG="$DATE - auto commit"
#================================================================================

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

# If no commit message is passed, then use the auto-commit message.
if [ -z $1 ]; then
    COMMIT_MSG=$GIT_AUTOCOMMIT_MSG
else
    if [ $1 == "-h" ]; then
        printMessage "${BLUE}HELP:${NC} Usage ${ORANGE}./POST.sh \"optional commit message\"${NC}"
        exit 0
    else
        COMMIT_MSG=$1
    fi
fi

printMessage "${BLUE}INFO:${NC} Starting post operation"

./godelw verify
printMessage "${BLUE}INFO:${NC} Verifying files"

handleError $? "Can't post due to errors. Please review" "${ORANGE}./godelw verify${NC}"

printMessage "${GREEN}SUCCESS:${NC} All checks passed"

printMessage "${BLUE}INFO:${NC} Generating auto-commit: ${ORANGE}${COMMIT_MSG}${NC}"
git add .
git commit --author="${GIT_USER}" -m "${COMMIT_MSG}"

printMessage "${BLUE}INFO:${NC} Pushing to Git"
git push $GIT_UPSTREAM $GIT_BRANCH
handleError $? "Can't post due to errors. Please review" "${ORANGE}git push ${GIT_UPSTREAM} ${GIT_BRANCH}${NC}"

 