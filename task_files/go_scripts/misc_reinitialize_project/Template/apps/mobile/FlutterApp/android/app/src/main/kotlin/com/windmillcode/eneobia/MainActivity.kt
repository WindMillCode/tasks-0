package com.[ORGANIZATION_NAME].[PROJECT_NAME]


// import android.content.pm.PackageInfo
// import android.content.pm.PackageManager
// import android.os.Bundle
// import android.util.Base64
// import android.util.Log

import io.flutter.embedding.android.FlutterActivity
// import io.flutter.embedding.engine.FlutterEngine
// import io.flutter.plugin.common.MethodChannel
// import java.security.MessageDigest
// import java.security.NoSuchAlgorithmException

class MainActivity : FlutterActivity()
// {
//   private val CHANNEL = "com.[ORGANIZATION_NAME].[PROJECT_NAME]/keyhash"

//   override fun configureFlutterEngine(flutterEngine: FlutterEngine) {
//     super.configureFlutterEngine(flutterEngine)
//     MethodChannel(flutterEngine.dartExecutor.binaryMessenger, CHANNEL).setMethodCallHandler { call, result ->
//       if (call.method == "getKeyHash") {
//         val keyHash = getKeyHash()
//         if (keyHash != null) {
//           result.success(keyHash)
//         } else {
//           result.error("UNAVAILABLE", "Key hash not available.", null)
//         }
//       } else {
//         result.notImplemented()
//       }
//     }
//   }

//   private fun getKeyHash(): String? {
//     try {
//       val info: PackageInfo = packageManager.getPackageInfo(
//         "com.[ORGANIZATION_NAME].[PROJECT_NAME]",
//         PackageManager.GET_SIGNATURES
//       )
//       for (signature in info.signatures) {
//         val md: MessageDigest = MessageDigest.getInstance("SHA")
//         md.update(signature.toByteArray())
//         return Base64.encodeToString(md.digest(), Base64.DEFAULT)
//       }
//     } catch (e: PackageManager.NameNotFoundException) {
//       e.printStackTrace()
//     } catch (e: NoSuchAlgorithmException) {
//       e.printStackTrace()
//     }
//     return null
//   }
// }
