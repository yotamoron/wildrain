
apps = {}

def addAicd(aicd):
    applicationName = aicd['applicationName']
    if not apps.has_key(applicationName):
        apps[applicationName] = {}
    application = apps[applicationName]
    version = aicd['version']
    if not application.has_key(version):
        application[version] = aicd

def getApplications():
    return apps
