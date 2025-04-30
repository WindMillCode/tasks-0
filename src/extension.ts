/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
import * as vscode from 'vscode';
import { WMLTasksJSONTaskProvider } from './TasksJSONProvider';

import { checkForGolang } from './checkForGolang';

let WMLDisposables: vscode.Disposable[] =[]
let WMLTaskProviders:any[] = [
	WMLTasksJSONTaskProvider
]
export async function activate(_context: vscode.ExtensionContext): Promise<void> {
	const workspaceRoot = (vscode.workspace.workspaceFolders && (vscode.workspace.workspaceFolders.length > 0))
		? vscode.workspace.workspaceFolders[0].uri.fsPath : undefined;


	let goInfo = await checkForGolang()
	if (!workspaceRoot) {
		return;
	}

  // TODO is this even being used like that
	WMLDisposables = WMLTaskProviders
	.map((providerType)=>{

		let providerInstance =  new providerType(workspaceRoot)
		providerInstance.goExecutable = goInfo.executable
		return vscode.tasks.registerTaskProvider(
			providerType.WindmillType,
			providerInstance
		);
	})

}

export function deactivate(): void {
	WMLDisposables
	.forEach((provider)=>{
		provider.dispose()
	})

}
